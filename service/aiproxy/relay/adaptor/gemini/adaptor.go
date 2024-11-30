package gemini

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"github.com/labring/sealos/service/aiproxy/common/config"
	"github.com/labring/sealos/service/aiproxy/common/helper"
	"github.com/labring/sealos/service/aiproxy/model"
	"github.com/labring/sealos/service/aiproxy/relay/adaptor/openai"
	"github.com/labring/sealos/service/aiproxy/relay/meta"
	relaymodel "github.com/labring/sealos/service/aiproxy/relay/model"
	"github.com/labring/sealos/service/aiproxy/relay/relaymode"
	"github.com/labring/sealos/service/aiproxy/relay/utils"
)

type Adaptor struct{}

const baseURL = "https://generativelanguage.googleapis.com"

func (a *Adaptor) GetRequestURL(meta *meta.Meta) (string, error) {
	version := helper.AssignOrDefault(meta.Channel.Config.APIVersion, config.GetGeminiVersion())
	var action string
	switch meta.Mode {
	case relaymode.Embeddings:
		action = "batchEmbedContents"
	default:
		action = "generateContent"
	}

	if meta.GetBool("stream") {
		action = "streamGenerateContent?alt=sse"
	}
	u := meta.Channel.BaseURL
	if u == "" {
		u = baseURL
	}
	return fmt.Sprintf("%s/%s/models/%s:%s", u, version, meta.ActualModelName, action), nil
}

func (a *Adaptor) SetupRequestHeader(meta *meta.Meta, c *gin.Context, req *http.Request) error {
	req.Header.Set("X-Goog-Api-Key", meta.Channel.Key)
	return nil
}

func (a *Adaptor) ConvertRequest(meta *meta.Meta, req *http.Request) (http.Header, io.Reader, error) {
	switch meta.Mode {
	case relaymode.Embeddings:
		request, err := utils.UnmarshalGeneralOpenAIRequest(req)
		if err != nil {
			return nil, nil, err
		}
		request.Model = meta.ActualModelName
		geminiEmbeddingRequest := ConvertEmbeddingRequest(request)
		data, err := json.Marshal(geminiEmbeddingRequest)
		if err != nil {
			return nil, nil, err
		}
		return nil, bytes.NewReader(data), nil
	case relaymode.ChatCompletions:
		data, err := ConvertRequest(meta, req)
		if err != nil {
			return nil, nil, err
		}
		data2, err := json.Marshal(data)
		if err != nil {
			return nil, nil, err
		}
		return nil, bytes.NewReader(data2), nil
	default:
		return nil, nil, errors.New("unsupported mode")
	}
}

func (a *Adaptor) DoRequest(meta *meta.Meta, c *gin.Context, req *http.Request) (*http.Response, error) {
	return utils.DoRequest(meta, c, req)
}

func (a *Adaptor) DoResponse(meta *meta.Meta, c *gin.Context, resp *http.Response) (usage *relaymodel.Usage, err *relaymodel.ErrorWithStatusCode) {
	switch meta.Mode {
	case relaymode.Embeddings:
		err, usage = EmbeddingHandler(c, resp)
	case relaymode.ChatCompletions:
		if utils.IsStreamResponse(resp) {
			var responseText string
			err, responseText = StreamHandler(c, resp)
			usage = openai.ResponseText2Usage(responseText, meta.ActualModelName, meta.PromptTokens)
		} else {
			err, usage = Handler(c, resp, meta.PromptTokens, meta.ActualModelName)
		}
	default:
		return nil, openai.ErrorWrapperWithMessage("unsupported mode", "unsupported_mode", http.StatusBadRequest)
	}
	return
}

func (a *Adaptor) GetModelList() []*model.ModelConfig {
	return ModelList
}

func (a *Adaptor) GetChannelName() string {
	return "google gemini"
}
