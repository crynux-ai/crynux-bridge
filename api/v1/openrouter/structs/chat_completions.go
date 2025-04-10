package structs

import "encoding/json"

type ChatCompletionsRequest struct {
	Messages           []CCReqMessage           `json:"messages" validate:"required"`
	Model              string                   `json:"model" validate:"required"`
	Audio              *CCReqAudio              `json:"audio"`
	FrequencyPenalty   float64                  `json:"frequency_penalty"`
	LogitBias          map[string]float64       `json:"logit_bias"`
	LogProbs           bool                     `json:"logprobs"`
	MaxComletionTokens int                      `json:"max_completion_tokens"`
	MetaData           map[string]string        `json:"metadata"`
	Modalities         []string                 `json:"modalities"`
	N                  int                      `json:"n"` // defaults to 1
	ParallelToolCalls  bool                     `json:"parallel_tool_calls"`
	Prediction         *CCReqPrediction         `json:"prediction"`
	PresencePenalty    float64                  `json:"presence_penalty,omitempty"`
	ReasoningEffort    string                   `json:"reasoning_effort"`
	ResponseFormat     json.RawMessage          `json:"response_format"`
	Seed               int                      `json:"seed"`
	ServiceTier        string                   `json:"service_tier"`
	Stop               json.RawMessage          `json:"stop"`
	Store              bool                     `json:"store"`
	Stream             bool                     `json:"stream"`
	StreamOptions      *CCReqStreamOptions      `json:"stream_options"`
	Temperature        float64                  `json:"temperature"`
	ToolChoice         json.RawMessage          `json:"tool_choice"`
	Tools              []map[string]interface{} `json:"tools"`
	TopLogprobs        int                      `json:"top_logprobs"`
	TopP               float64                  `json:"top_p"`
	User               string                   `json:"user"`
	WebSearchOptions   json.RawMessage          `json:"web_search_options"`
}

func (ccr *ChatCompletionsRequest) SetDefaultValues() {
	if ccr.N == 0 {
		ccr.N = 1 // defaults to 1
	}
}

// Chat Completions Request Message
type CCReqMessage struct {
	Role       ChatCompletionsRole    `json:"role" validate:"required"`
	Content    json.RawMessage        `json:"content"`
	Name       string                 `json:"name"`
	Audio      *CCReqMessageAudio      `json:"audio"`
	Refusal    string                 `json:"refusal"`
	ToolCalls  []CCReqMessageToolCall `json:"tool_calls"`
	ToolCallID string                 `json:"tool_call_id"`
}

type ChatCompletionsRole string

const (
	ChatCompletionsRoleDeveloper ChatCompletionsRole = "developer"
	ChatCompletionsRoleSystem    ChatCompletionsRole = "system"
	ChatCompletionsRoleUser      ChatCompletionsRole = "user"
	ChatCompletionsRoleAssistant ChatCompletionsRole = "assistant"
	ChatCompletionsRoleTool      ChatCompletionsRole = "tool"
	ChatCompletionsRoleUnknown   ChatCompletionsRole = "unknown role"
)

type CCReqMessageAudio struct {
	ID string `json:"id" validate:"required"`
}

type CCReqMessageToolCall struct {
	ID       string                       `json:"id" validate:"required"`
	Function CCReqMessageToolCallFunction `json:"function" validate:"required"`
	Type     string                       `json:"type" validate:"required"`
}

type CCReqMessageToolCallFunction struct {
	Name      string `json:"name" validate:"required"`
	Arguments string `json:"arguments" validate:"required"`
}

type CCReqPrediction struct {
	StaticContent StaticContent
}

type StaticContent struct {
	Content json.RawMessage `json:"content" validate:"required"`
	Type    string          `json:"type" validate:"required"`
}

type CCReqStreamOptions struct {
	IncludeUsage bool `json:"include_usage"`
}

type CCReqTool struct {
	Function json.RawMessage `json:"function"`
	Type     string          `json:"type"`
}

type CCReqAudio struct {
	Format string `json:"format" validate:"required"`
	Voice  string `json:"voice" validate:"required"`
}

type ChatCompletionsResponse struct {
	Id          string        `json:"id"`
	Object      string        `json:"object"`
	Created     int64         `json:"created"`
	Model       string        `json:"model"`
	Choices     []CCResChoice `json:"choices"`
	Usage       CCResUsage    `json:"usage"`
	ServiceTier string        `json:"service_tier"`
}

type CCResChoice struct {
	Index        int          `json:"index"`
	Message      CCResMessage `json:"message"`
	LogProbs     string       `json:"logprobs"`
	FinishReason string       `json:"finish_reason"`
}

type CCResMessage struct {
	Role        ChatCompletionsRole      `json:"role"`
	Content     string                   `json:"content"`
	Refusal     string                   `json:"refusal"`
	Annotations []interface{}            `json:"annotations"`
	Audio       interface{}              `json:"audio"`
	ToolCalls   []map[string]interface{} `json:"tool_calls"`
}

type CCResUsage struct {
	PromptTokens            int                     `json:"prompt_tokens"`
	CompletionTokens        int                     `json:"completion_tokens"`
	TotalTokens             int                     `json:"total_tokens"`
	PromptTokensDetails     PromptTokensDetails     `json:"prompt_tokens_details"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}

type PromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
	AudioTokens  int `json:"audio_tokens"`
}

type CompletionTokensDetails struct {
	ReasoningTokens          int `json:"reasoning_tokens"`
	AudioTokens              int `json:"audio_tokens"`
	AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
	RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
}
