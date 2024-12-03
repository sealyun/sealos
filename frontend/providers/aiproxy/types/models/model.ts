export enum ModelType {
  OpenAI = '1',
  API2D = '2',
  Azure = '3',
  CloseAI = '4',
  OpenAISB = '5',
  OpenAIMax = '6',
  OhMyGPT = '7',
  Custom = '8',
  Ails = '9',
  AIProxy = '10',
  PaLM = '11',
  API2GPT = '12',
  AIGC2D = '13',
  Anthropic = '14',
  Baidu = '15',
  Zhipu = '16',
  Ali = '17',
  Xunfei = '18',
  AI360 = '19',
  OpenRouter = '20',
  AIProxyLibrary = '21',
  FastGPT = '22',
  Tencent = '23',
  Gemini = '24',
  Moonshot = '25',
  Baichuan = '26',
  Minimax = '27',
  Mistral = '28',
  Groq = '29',
  Ollama = '30',
  LingYiWanWu = '31',
  StepFun = '32',
  AwsClaude = '33',
  Coze = '34',
  Cohere = '35',
  DeepSeek = '36',
  Cloudflare = '37',
  DeepL = '38',
  TogetherAI = '39',
  Doubao = '40',
  Novita = '41',
  VertextAI = '42',
  SiliconFlow = '43'
}

export type ModelMap = { [K in ModelType]?: string[] }

export type ModelMappingMap = { [K in ModelType]?: {} }

export interface ModelConfig {
  image_prices: null
  model: string
  owner: string
  image_batch_size: number
  type: number
  input_price: number
  output_price: number
  created_at: number
  updated_at: number
}
