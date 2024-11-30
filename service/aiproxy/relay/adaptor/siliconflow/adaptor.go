package siliconflow

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labring/sealos/service/aiproxy/common/client"
	"github.com/labring/sealos/service/aiproxy/model"
	"github.com/labring/sealos/service/aiproxy/relay/adaptor"
	"github.com/labring/sealos/service/aiproxy/relay/adaptor/openai"
	"github.com/labring/sealos/service/aiproxy/relay/meta"
)

var (
	_ adaptor.Adaptor    = (*Adaptor)(nil)
	_ adaptor.GetBalance = (*Adaptor)(nil)
)

type Adaptor struct {
	openai.Adaptor
}

const baseURL = "https://api.siliconflow.cn"

func (a *Adaptor) GetRequestURL(meta *meta.Meta) (string, error) {
	if meta.Channel.BaseURL == "" {
		meta.Channel.BaseURL = baseURL
	}
	return a.Adaptor.GetRequestURL(meta)
}

func (a *Adaptor) GetModelList() []*model.ModelConfig {
	return ModelList
}

func (a *Adaptor) GetChannelName() string {
	return "siliconflow"
}

func (a *Adaptor) GetBalance(channel *model.Channel) (float64, error) {
	url := "https://api.siliconflow.cn/v1/user/info"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Bearer "+channel.Key)
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status code: %d", res.StatusCode)
	}
	response := UsageResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return 0, err
	}
	balance, err := strconv.ParseFloat(response.Data.Balance, 64)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

type UsageResponse struct {
	Message string `json:"message"`
	Data    struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Image         string `json:"image"`
		Email         string `json:"email"`
		Balance       string `json:"balance"`
		Status        string `json:"status"`
		Introduction  string `json:"introduction"`
		Role          string `json:"role"`
		ChargeBalance string `json:"chargeBalance"`
		TotalBalance  string `json:"totalBalance"`
		Category      string `json:"category"`
		IsAdmin       bool   `json:"isAdmin"`
	} `json:"data"`
	Code   int  `json:"code"`
	Status bool `json:"status"`
}
