package structs

/* Request */

type CompletionsRequest struct {
	Model             string             `json:"model" validate:"required"`
	Prompt            string             `json:"prompt" validate:"required"`
	Stream            bool               `json:"stream"` // default to false
	MaxTokens         *int               `json:"max_tokens"`
	Temperature       float64            `json:"temperature"`
	Seed              int                `json:"seed"`
	TopP              *float64           `json:"top_p"`
	TopK              *int               `json:"top_k"`
	FrequencyPenalty  *float64           `json:"frequency_penalty"`
	PresencePenalty   *float64           `json:"presence_penalty"`
	RepetitionPenalty *float64           `json:"repetition_penalty"`
	LogitBias         map[string]float64 `json:"logit_bias"`
	TopLogprobs       int                `json:"top_logprobs"`
	MinP              *float64           `json:"min_p"`
	TopA              *float64           `json:"top_a"`

	// Transforms []string `json:"transforms"` // openrouter only
	// Models     []string `json:"models"`
	// Provider   TODO     `json:"provider"`
	// Reasoning  TODO     `json:"reasoning"`

	Stop []string `json:"stop"`

	BestOf        int               `json:"best_of"` // defaults to 1
	Echo          bool              `json:"echo"`
	LogProbs      int               `json:"logprobs"`
	N             int               `json:"n" default:"1"`
	StreamOptions CReqStreamOptions `json:"stream_options"`
	Suffix        string            `json:"suffix"`
	User          string            `json:"user"`
}

func (cr *CompletionsRequest) SetDefaultValues() {
	if cr.BestOf == 0 {
		cr.BestOf = 1
	}
}

type CReqStreamOptions struct {
	IncludeUsage bool `json:"include_usage"`
}

/* Response */

type CompletionsResponse struct {
	Id                string       `json:"id"`
	Object            string       `json:"object"`
	Created           int64        `json:"created"`
	Model             string       `json:"model"`
	SystemFingerprint string       `json:"system_fingerprint"`
	Choices           []CResChoice `json:"choices"`
	Usage             CResUsage    `json:"usage"`
}

type CResChoice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	LogProbs     string `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type CResUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
