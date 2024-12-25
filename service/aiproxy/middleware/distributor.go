package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labring/sealos/service/aiproxy/common/config"
	"github.com/labring/sealos/service/aiproxy/common/ctxkey"
	"github.com/labring/sealos/service/aiproxy/model"
	"github.com/labring/sealos/service/aiproxy/relay/meta"
	"github.com/labring/sealos/service/aiproxy/relay/relaymode"
)

const (
	groupModelRPMKey = "group_model_rpm:%s:%s"
)

type ModelRequest struct {
	Model string `form:"model" json:"model"`
}

func calculateGroupConsumeLevelRpmRatio(usedAmount float64) float64 {
	v := config.GetGroupConsumeLevelRpmRatio()
	var maxConsumeLevel float64 = -1
	var groupConsumeLevelRpmRatio float64
	for consumeLevel, ratio := range v {
		if usedAmount < consumeLevel {
			continue
		}
		if consumeLevel > maxConsumeLevel {
			maxConsumeLevel = consumeLevel
			groupConsumeLevelRpmRatio = ratio
		}
	}
	if groupConsumeLevelRpmRatio <= 0 {
		groupConsumeLevelRpmRatio = 1
	}
	return groupConsumeLevelRpmRatio
}

func getGroupRPMRatio(group *model.GroupCache) float64 {
	groupRPMRatio := group.RPMRatio
	if groupRPMRatio <= 0 {
		groupRPMRatio = 1
	}
	return groupRPMRatio
}

func checkModelRPM(c *gin.Context, group *model.GroupCache, requestModel string, modelRPM int64) bool {
	if modelRPM <= 0 {
		return true
	}

	groupConsumeLevelRpmRatio := calculateGroupConsumeLevelRpmRatio(group.UsedAmount)
	groupRPMRatio := getGroupRPMRatio(group)

	adjustedModelRPM := int64(float64(modelRPM) * groupRPMRatio * groupConsumeLevelRpmRatio)

	ok := ForceRateLimit(
		c.Request.Context(),
		fmt.Sprintf(groupModelRPMKey, group.ID, requestModel),
		adjustedModelRPM,
		time.Minute,
	)

	if !ok {
		abortWithMessage(c, http.StatusTooManyRequests,
			group.ID+" is requesting too frequently",
		)
		return false
	}
	return true
}

func Distribute(c *gin.Context) {
	if config.GetDisableServe() {
		abortWithMessage(c, http.StatusServiceUnavailable, "service is under maintenance")
		return
	}

	log := GetLogger(c)

	group := c.MustGet(ctxkey.Group).(*model.GroupCache)

	requestModel, err := getRequestModel(c)
	if err != nil {
		abortWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	if requestModel == "" {
		abortWithMessage(c, http.StatusBadRequest, "no model provided")
		return
	}

	SetLogModelFields(log.Data, requestModel)

	token := c.MustGet(ctxkey.Token).(*model.TokenCache)
	if len(token.Models) == 0 || !slices.Contains(token.Models, requestModel) {
		abortWithMessage(c,
			http.StatusForbidden,
			fmt.Sprintf("token (%s[%d]) has no permission to use model: %s",
				token.Name, token.ID, requestModel,
			),
		)
		return
	}

	mc, ok := model.CacheGetModelConfig(requestModel)
	if !ok {
		abortWithMessage(c, http.StatusServiceUnavailable, requestModel+" is not available")
		return
	}

	if !checkModelRPM(c, group, requestModel, mc.RPM) {
		return
	}

	c.Set(ctxkey.OriginalModel, requestModel)

	c.Next()
}

func NewMetaByContext(c *gin.Context, channel *model.Channel) *meta.Meta {
	originalModel := c.MustGet(ctxkey.OriginalModel).(string)
	requestID := c.GetString(ctxkey.RequestID)
	group := c.MustGet(ctxkey.Group).(*model.GroupCache)
	token := c.MustGet(ctxkey.Token).(*model.TokenCache)
	return meta.NewMeta(
		channel,
		relaymode.GetByPath(c.Request.URL.Path),
		originalModel,
		meta.WithRequestID(requestID),
		meta.WithGroup(group),
		meta.WithToken(token),
	)
}
