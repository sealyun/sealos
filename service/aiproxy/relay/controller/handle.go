package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labring/sealos/service/aiproxy/common"
	"github.com/labring/sealos/service/aiproxy/common/config"
	"github.com/labring/sealos/service/aiproxy/common/conv"
	"github.com/labring/sealos/service/aiproxy/middleware"
	"github.com/labring/sealos/service/aiproxy/model"
	"github.com/labring/sealos/service/aiproxy/relay/adaptor/openai"
	"github.com/labring/sealos/service/aiproxy/relay/channeltype"
	"github.com/labring/sealos/service/aiproxy/relay/meta"
	relaymodel "github.com/labring/sealos/service/aiproxy/relay/model"
)

func Handle(meta *meta.Meta, c *gin.Context, preProcess func() (*PreCheckGroupBalanceReq, error)) *relaymodel.ErrorWithStatusCode {
	log := middleware.GetLogger(c)
	ctx := c.Request.Context()

	// 1. Get adaptor
	adaptor, ok := channeltype.GetAdaptor(meta.Channel.Type)
	if !ok {
		log.Errorf("invalid (%s[%d]) channel type: %d", meta.Channel.Name, meta.Channel.ID, meta.Channel.Type)
		return openai.ErrorWrapperWithMessage("invalid channel error", "invalid_channel_type", http.StatusInternalServerError)
	}

	// 2. Get group balance
	groupRemainBalance, postGroupConsumer, err := getGroupBalance(ctx, meta)
	if err != nil {
		log.Errorf("get group (%s) balance failed: %v", meta.Group.ID, err)
		return openai.ErrorWrapper(
			fmt.Errorf("get group (%s) balance failed", meta.Group.ID),
			"get_group_quota_failed",
			http.StatusInternalServerError,
		)
	}

	// 3. Pre-process request
	preCheckReq, err := preProcess()
	if err != nil {
		log.Errorf("pre-process request failed: %s", err.Error())
		var detail *model.RequestDetail
		body, bodyErr := common.GetRequestBody(c.Request)
		if bodyErr != nil {
			log.Errorf("get request body failed: %s", bodyErr.Error())
		} else {
			detail = &model.RequestDetail{
				RequestBody: conv.BytesToString(body),
			}
		}
		ConsumeWaitGroup.Add(1)
		go postConsumeAmount(context.Background(),
			&ConsumeWaitGroup,
			nil,
			http.StatusBadRequest,
			nil,
			meta,
			0,
			0,
			err.Error(),
			detail,
		)
		return openai.ErrorWrapper(err, "invalid_request", http.StatusBadRequest)
	}

	// 4. Pre-check balance
	ok = checkGroupBalance(preCheckReq, meta, groupRemainBalance)
	if !ok {
		return openai.ErrorWrapper(errors.New("group balance is not enough"), "insufficient_group_balance", http.StatusForbidden)
	}

	meta.InputTokens = preCheckReq.InputTokens

	// 5. Do request
	usage, detail, respErr := DoHelper(adaptor, c, meta)
	if respErr != nil {
		if detail != nil && config.DebugEnabled {
			log.Errorf(
				"handle failed: %+v\nrequest detail:\n%s\nresponse detail:\n%s",
				respErr.Error,
				detail.RequestBody,
				detail.ResponseBody,
			)
		} else {
			log.Errorf("handle failed: %+v", respErr.Error)
		}

		ConsumeWaitGroup.Add(1)
		go postConsumeAmount(context.Background(),
			&ConsumeWaitGroup,
			postGroupConsumer,
			respErr.StatusCode,
			usage,
			meta,
			preCheckReq.InputPrice,
			preCheckReq.OutputPrice,
			respErr.Error.String(),
			detail,
		)
		return respErr
	}

	// 6. Post consume
	ConsumeWaitGroup.Add(1)
	go postConsumeAmount(context.Background(),
		&ConsumeWaitGroup,
		postGroupConsumer,
		http.StatusOK,
		usage,
		meta,
		preCheckReq.InputPrice,
		preCheckReq.OutputPrice,
		"",
		nil,
	)

	return nil
}
