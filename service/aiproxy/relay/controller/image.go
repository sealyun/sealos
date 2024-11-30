package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labring/sealos/service/aiproxy/common/logger"
	"github.com/labring/sealos/service/aiproxy/relay/adaptor/openai"
	"github.com/labring/sealos/service/aiproxy/relay/channeltype"
	"github.com/labring/sealos/service/aiproxy/relay/meta"
	relaymodel "github.com/labring/sealos/service/aiproxy/relay/model"
	billingprice "github.com/labring/sealos/service/aiproxy/relay/price"
	"github.com/labring/sealos/service/aiproxy/relay/utils"
)

func getImageRequest(c *gin.Context) (*relaymodel.ImageRequest, error) {
	imageRequest, err := utils.UnmarshalImageRequest(c.Request)
	if err != nil {
		return nil, err
	}
	if imageRequest.N == 0 {
		imageRequest.N = 1
	}
	if imageRequest.Size == "" {
		imageRequest.Size = "1024x1024"
	}
	return imageRequest, nil
}

func validateImageRequest(imageRequest *relaymodel.ImageRequest) *relaymodel.ErrorWithStatusCode {
	// check prompt length
	if imageRequest.Prompt == "" {
		return openai.ErrorWrapper(errors.New("prompt is required"), "prompt_missing", http.StatusBadRequest)
	}

	// Number of generated images validation
	if !billingprice.ValidateImageMaxBatchSize(imageRequest.Model, imageRequest.N) {
		return openai.ErrorWrapper(errors.New("invalid value of n"), "n_not_within_range", http.StatusBadRequest)
	}
	return nil
}

func getImageCostPrice(modelName string, reqModel string, size string) (float64, error) {
	imageCostPrice, ok := billingprice.GetImageSizePrice(modelName, reqModel, size)
	if !ok {
		return 0, fmt.Errorf("invalid image size: %s", size)
	}
	return imageCostPrice, nil
}

func RelayImageHelper(c *gin.Context, _ int) *relaymodel.ErrorWithStatusCode {
	ctx := c.Request.Context()
	imageRequest, err := getImageRequest(c)
	if err != nil {
		logger.Errorf(ctx, "getImageRequest failed: %s", err.Error())
		return openai.ErrorWrapper(err, "invalid_image_request", http.StatusBadRequest)
	}

	meta := meta.GetByContext(c)

	meta.PromptTokens = imageRequest.N

	bizErr := validateImageRequest(imageRequest)
	if bizErr != nil {
		return bizErr
	}

	imageCostPrice, err := getImageCostPrice(meta.OriginModelName, meta.ActualModelName, imageRequest.Size)
	if err != nil {
		return openai.ErrorWrapper(err, "get_image_cost_price_failed", http.StatusInternalServerError)
	}

	c.Set("response_format", imageRequest.ResponseFormat)

	adaptor, ok := channeltype.GetAdaptor(meta.Channel.Type)
	if !ok {
		return openai.ErrorWrapper(fmt.Errorf("invalid channel type: %d", meta.Channel.Type), "invalid_channel_type", http.StatusBadRequest)
	}

	ok, postGroupConsumer, err := preCheckGroupBalance(ctx, &PreCheckGroupBalanceReq{
		PromptTokens: meta.PromptTokens,
		Price:        imageCostPrice,
	}, meta)
	if err != nil {
		logger.Errorf(ctx, "get group (%s) balance failed: %v", meta.Group.ID, err)
		return openai.ErrorWrapper(
			fmt.Errorf("get group (%s) balance failed", meta.Group.ID),
			"get_group_quota_failed",
			http.StatusInternalServerError,
		)
	}
	if !ok {
		return openai.ErrorWrapper(errors.New("group balance is not enough"), "insufficient_group_balance", http.StatusForbidden)
	}

	// do response
	usage, respErr := DoHelper(adaptor, c, meta)
	if respErr != nil {
		logger.Errorf(ctx, "do response failed: %s", respErr)
		ConsumeWaitGroup.Add(1)
		go postConsumeAmount(context.Background(),
			&ConsumeWaitGroup,
			postGroupConsumer,
			respErr.StatusCode,
			c.Request.URL.Path,
			usage,
			meta,
			imageCostPrice,
			0,
			respErr.String(),
		)
		return respErr
	}

	ConsumeWaitGroup.Add(1)
	go postConsumeAmount(context.Background(),
		&ConsumeWaitGroup,
		postGroupConsumer,
		http.StatusOK,
		c.Request.URL.Path,
		usage, meta, imageCostPrice, 0, imageRequest.Size,
	)

	return nil
}
