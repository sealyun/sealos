package controller

import (
	"fmt"
	"net/http"
	"slices"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"github.com/labring/sealos/service/aiproxy/common/config"
	"github.com/labring/sealos/service/aiproxy/common/ctxkey"
	"github.com/labring/sealos/service/aiproxy/model"
	"github.com/labring/sealos/service/aiproxy/relay/channeltype"
	relaymodel "github.com/labring/sealos/service/aiproxy/relay/model"
)

// https://platform.openai.com/docs/api-reference/models/list

type OpenAIModelPermission struct {
	Group              *string `json:"group"`
	ID                 string  `json:"id"`
	Object             string  `json:"object"`
	Organization       string  `json:"organization"`
	Created            int     `json:"created"`
	AllowCreateEngine  bool    `json:"allow_create_engine"`
	AllowSampling      bool    `json:"allow_sampling"`
	AllowLogprobs      bool    `json:"allow_logprobs"`
	AllowSearchIndices bool    `json:"allow_search_indices"`
	AllowView          bool    `json:"allow_view"`
	AllowFineTuning    bool    `json:"allow_fine_tuning"`
	IsBlocking         bool    `json:"is_blocking"`
}

type OpenAIModels struct {
	Parent     *string                 `json:"parent"`
	ID         string                  `json:"id"`
	Object     string                  `json:"object"`
	OwnedBy    string                  `json:"owned_by"`
	Root       string                  `json:"root"`
	Permission []OpenAIModelPermission `json:"permission"`
	Created    int                     `json:"created"`
}

type BuiltinModelConfig model.ModelConfig

func (c *BuiltinModelConfig) MarshalJSON() ([]byte, error) {
	type Alias BuiltinModelConfig
	return json.Marshal(&struct {
		*Alias
		CreatedAt int64 `json:"created_at,omitempty"`
		UpdatedAt int64 `json:"updated_at,omitempty"`
	}{
		Alias: (*Alias)(c),
	})
}

var (
	models                  []OpenAIModels
	modelsMap               map[string]OpenAIModels
	builtinChannelID2Models map[int][]*BuiltinModelConfig
)

func init() {
	var permission []OpenAIModelPermission
	permission = append(permission, OpenAIModelPermission{
		ID:                 "modelperm-LwHkVFn8AcMItP432fKKDIKJ",
		Object:             "model_permission",
		Created:            1626777600,
		AllowCreateEngine:  true,
		AllowSampling:      true,
		AllowLogprobs:      true,
		AllowSearchIndices: false,
		AllowView:          true,
		AllowFineTuning:    false,
		Organization:       "*",
		Group:              nil,
		IsBlocking:         false,
	})

	builtinChannelID2Models = make(map[int][]*BuiltinModelConfig)
	// https://platform.openai.com/docs/models/model-endpoint-compatibility
	for i, adaptor := range channeltype.ChannelAdaptor {
		channelName := adaptor.GetChannelName()
		modelNames := adaptor.GetModelList()
		for _, model := range modelNames {
			models = append(models, OpenAIModels{
				ID:         model.Model,
				Object:     "model",
				Created:    1626777600,
				OwnedBy:    channelName,
				Permission: permission,
				Root:       model.Model,
				Parent:     nil,
			})
		}
		builtinChannelID2Models[i] = make([]*BuiltinModelConfig, len(modelNames))
		for idx, model := range modelNames {
			builtinChannelID2Models[i][idx] = (*BuiltinModelConfig)(model)
		}
	}
	modelsMap = make(map[string]OpenAIModels)
	for _, model := range models {
		modelsMap[model.ID] = model
	}
	for _, models := range builtinChannelID2Models {
		sort.Slice(models, func(i, j int) bool {
			return models[i].Model < models[j].Model
		})
	}
}

func BuiltinModels(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    builtinChannelID2Models,
	})
}

func BuiltinModelsByType(c *gin.Context) {
	channelType := c.Param("type")
	if channelType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "type is required",
		})
		return
	}
	channelTypeInt, err := strconv.Atoi(channelType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid type",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    builtinChannelID2Models[channelTypeInt],
	})
}

func ChannelDefaultModelsAndMapping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"models":  config.GetDefaultChannelModels(),
			"mapping": config.GetDefaultChannelModelMapping(),
		},
	})
}

func ChannelDefaultModelsAndMappingByType(c *gin.Context) {
	channelType := c.Param("type")
	if channelType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "type is required",
		})
		return
	}
	channelTypeInt, err := strconv.Atoi(channelType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid type",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"models":  config.GetDefaultChannelModels()[channelTypeInt],
			"mapping": config.GetDefaultChannelModelMapping()[channelTypeInt],
		},
	})
}

func EnabledModels(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    model.CacheGetAllModelsAndConfig(),
	})
}

func ChannelEnabledModels(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    model.CacheGetAllChannelModelsAndConfig(),
	})
}

func ChannelEnabledModelsByType(c *gin.Context) {
	channelTypeStr := c.Param("type")
	if channelTypeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "type is required",
		})
		return
	}
	channelTypeInt, err := strconv.Atoi(channelTypeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid type",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    model.CacheGetAllChannelModelsAndConfig()[channelTypeInt],
	})
}

func ListModels(c *gin.Context) {
	channel := c.MustGet(ctxkey.Channel).(*model.Channel)
	availableOpenAIModels := make([]OpenAIModels, 0, len(channel.Models))

	for _, modelName := range channel.Models {
		if model, ok := modelsMap[modelName]; ok {
			availableOpenAIModels = append(availableOpenAIModels, model)
			continue
		}
		availableOpenAIModels = append(availableOpenAIModels, OpenAIModels{
			ID:      modelName,
			Object:  "model",
			Created: 1626777600,
			OwnedBy: "custom",
			Root:    modelName,
			Parent:  nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"object": "list",
		"data":   availableOpenAIModels,
	})
}

func RetrieveModel(c *gin.Context) {
	channel := c.MustGet(ctxkey.Channel).(*model.Channel)
	modelID := c.Param("model")
	model, ok := modelsMap[modelID]
	if !ok || !slices.Contains(channel.Models, modelID) {
		c.JSON(200, gin.H{
			"error": relaymodel.Error{
				Message: fmt.Sprintf("the model '%s' does not exist", modelID),
				Type:    "invalid_request_error",
				Param:   "model",
				Code:    "model_not_found",
			},
		})
		return
	}
	c.JSON(200, model)
}
