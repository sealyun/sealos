package ali

import (
	"github.com/labring/sealos/service/aiproxy/model"
	"github.com/labring/sealos/service/aiproxy/relay/relaymode"
)

// https://help.aliyun.com/zh/model-studio/getting-started/models?spm=a2c4g.11186623.0.i12#ced16cb6cdfsy

var ModelList = []*model.ModelConfig{
	// 通义千问-Max
	{
		Model:       "qwen-max",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.02,
		OutputPrice: 0.06,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(30720),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-max-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.02,
		OutputPrice: 0.06,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32768),
			model.WithModelConfigMaxInputTokens(30720),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问-Plus
	{
		Model:       "qwen-plus",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0008,
		OutputPrice: 0.002,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-plus-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0008,
		OutputPrice: 0.002,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(8000),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问-Turbo
	{
		Model:       "qwen-turbo",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0003,
		OutputPrice: 0.0006,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-turbo-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0003,
		OutputPrice: 0.0006,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(1000000),
			model.WithModelConfigMaxInputTokens(1000000),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	// Qwen-Long
	{
		Model:       "qwen-long",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0005,
		OutputPrice: 0.002,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(10000000),
			model.WithModelConfigMaxInputTokens(10000000),
			model.WithModelConfigMaxOutputTokens(6000),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问VL
	{
		Model:       "qwen-vl-max",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.02,
		OutputPrice: 0.02,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-vl-max-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.02,
		OutputPrice: 0.02,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-vl-plus",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.008,
		OutputPrice: 0.008,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8000),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-vl-plus-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.008,
		OutputPrice: 0.008,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigVision(true),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问OCR
	{
		Model:       "qwen-vl-ocr",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.005,
		OutputPrice: 0.005,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(34096),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigVision(true),
		),
	},
	{
		Model:       "qwen-vl-ocr-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.005,
		OutputPrice: 0.005,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(34096),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(4096),
			model.WithModelConfigVision(true),
		),
	},

	// 通义千问Math
	{
		Model:       "qwen-math-plus",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.004,
		OutputPrice: 0.012,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-math-plus-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.004,
		OutputPrice: 0.012,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-math-turbo",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.002,
		OutputPrice: 0.006,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-math-turbo-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.002,
		OutputPrice: 0.006,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问Coder
	{
		Model:       "qwen-coder-plus",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0035,
		OutputPrice: 0.007,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-coder-plus-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0035,
		OutputPrice: 0.007,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-coder-turbo",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.002,
		OutputPrice: 0.006,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-coder-turbo-latest",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.002,
		OutputPrice: 0.006,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问2.5
	{
		Model:       "qwen2.5-72b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.004,
		OutputPrice: 0.012,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2.5-32b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0035,
		OutputPrice: 0.007,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2.5-14b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.002,
		OutputPrice: 0.006,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2.5-7b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.001,
		OutputPrice: 0.002,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问2
	{
		Model:       "qwen2-72b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.004,
		OutputPrice: 0.012,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(128000),
			model.WithModelConfigMaxOutputTokens(6144),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2-57b-a14b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0035,
		OutputPrice: 0.007,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(65536),
			model.WithModelConfigMaxInputTokens(63488),
			model.WithModelConfigMaxOutputTokens(6144),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2-7b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.001,
		OutputPrice: 0.002,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(128000),
			model.WithModelConfigMaxOutputTokens(6144),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问1.5
	{
		Model:       "qwen1.5-110b-chat",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.007,
		OutputPrice: 0.014,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(8000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen1.5-72b-chat",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.005,
		OutputPrice: 0.01,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(8000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen1.5-32b-chat",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0035,
		OutputPrice: 0.007,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(8000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen1.5-14b-chat",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.002,
		OutputPrice: 0.004,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8000),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen1.5-7b-chat",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.001,
		OutputPrice: 0.002,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8000),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问
	{
		Model:       "qwen-72b-chat",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.02,
		OutputPrice: 0.02,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(32000),
			model.WithModelConfigMaxInputTokens(30000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-14b-chat",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.008,
		OutputPrice: 0.008,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(8000),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(2000),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen-7b-chat",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.006,
		OutputPrice: 0.006,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(7500),
			model.WithModelConfigMaxInputTokens(6000),
			model.WithModelConfigMaxOutputTokens(1500),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问数学模型
	{
		Model:       "qwen2.5-math-72b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.004,
		OutputPrice: 0.012,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2.5-math-7b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.001,
		OutputPrice: 0.002,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2-math-72b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.004,
		OutputPrice: 0.012,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2-math-7b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.001,
		OutputPrice: 0.002,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4096),
			model.WithModelConfigMaxInputTokens(3072),
			model.WithModelConfigMaxOutputTokens(3072),
			model.WithModelConfigToolChoice(true),
		),
	},

	// 通义千问Coder
	{
		Model:       "qwen2.5-coder-32b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.0035,
		OutputPrice: 0.007,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2.5-coder-14b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.002,
		OutputPrice: 0.006,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},
	{
		Model:       "qwen2.5-coder-7b-instruct",
		Type:        relaymode.ChatCompletions,
		Owner:       model.ModelOwnerAlibaba,
		InputPrice:  0.001,
		OutputPrice: 0.002,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(131072),
			model.WithModelConfigMaxInputTokens(129024),
			model.WithModelConfigMaxOutputTokens(8192),
			model.WithModelConfigToolChoice(true),
		),
	},

	// stable-diffusion
	{
		Model: "stable-diffusion-xl",
		Type:  relaymode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
	},
	{
		Model: "stable-diffusion-v1.5",
		Type:  relaymode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
	},
	{
		Model: "stable-diffusion-3.5-large",
		Type:  relaymode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
	},
	{
		Model: "stable-diffusion-3.5-large-turbo",
		Type:  relaymode.ImagesGenerations,
		Owner: model.ModelOwnerStabilityAI,
	},
	{
		Model:      "sambert-v1",
		Type:       relaymode.AudioSpeech,
		Owner:      model.ModelOwnerAlibaba,
		InputPrice: 0.1,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(10000),
			model.WithModelConfigSupportFormats([]string{"mp3", "wav", "pcm"}),
			model.WithModelConfigSupportVoices([]string{
				"zhinan",
				"zhiqi",
				"zhichu",
				"zhide",
				"zhijia",
				"zhiru",
				"zhiqian",
				"zhixiang",
				"zhiwei",
				"zhihao",
				"zhijing",
				"zhiming",
				"zhimo",
				"zhina",
				"zhishu",
				"zhistella",
				"zhiting",
				"zhixiao",
				"zhiya",
				"zhiye",
				"zhiying",
				"zhiyuan",
				"zhiyue",
				"zhigui",
				"zhishuo",
				"zhimiao-emo",
				"zhimao",
				"zhilun",
				"zhifei",
				"zhida",
				"indah",
				"clara",
				"hanna",
				"beth",
				"betty",
				"cally",
				"cindy",
				"eva",
				"donna",
				"brian",
				"waan",
			}),
		),
	},

	{
		Model: "paraformer-realtime-v2",
		Type:  relaymode.AudioTranscription,
		Owner: model.ModelOwnerAlibaba,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(10000),
			model.WithModelConfigSupportFormats([]string{"pcm", "wav", "opus", "speex", "aac", "amr"}),
		),
	},

	{
		Model: "gte-rerank",
		Type:  relaymode.Rerank,
		Owner: model.ModelOwnerAlibaba,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxContextTokens(4000),
			model.WithModelConfigMaxInputTokens(4000),
		),
	},

	{
		Model:      "text-embedding-v1",
		Type:       relaymode.Embeddings,
		Owner:      model.ModelOwnerAlibaba,
		InputPrice: 0.0007,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(2048),
		),
	},
	{
		Model:      "text-embedding-v2",
		Type:       relaymode.Embeddings,
		Owner:      model.ModelOwnerAlibaba,
		InputPrice: 0.0007,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(2048),
		),
	},
	{
		Model:      "text-embedding-v3",
		Type:       relaymode.Embeddings,
		Owner:      model.ModelOwnerAlibaba,
		InputPrice: 0.0007,
		Config: model.NewModelConfig(
			model.WithModelConfigMaxInputTokens(8192),
		),
	},
}
