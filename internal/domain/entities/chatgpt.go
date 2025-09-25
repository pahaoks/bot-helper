package entities

type ChatGPTRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
	Store bool   `json:"store"`
}

type ChatGPTResponse struct {
	ID         string `json:"id"`
	Object     string `json:"object"`
	CreatedAt  int    `json:"created_at"`
	Status     string `json:"status"`
	Background bool   `json:"background"`
	Billing    struct {
		Payer string `json:"payer"`
	} `json:"billing"`
	Error *struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Param   interface{} `json:"param"`
		Type    string      `json:"type"`
	} `json:"error"`
	IncompleteDetails interface{} `json:"incomplete_details"`
	Instructions      interface{} `json:"instructions"`
	MaxOutputTokens   interface{} `json:"max_output_tokens"`
	MaxToolCalls      interface{} `json:"max_tool_calls"`
	Model             string      `json:"model"`
	Output            []struct {
		ID      string        `json:"id"`
		Type    string        `json:"type"`
		Summary []interface{} `json:"summary,omitempty"`
		Status  string        `json:"status,omitempty"`
		Content []struct {
			Type        string        `json:"type"`
			Annotations []interface{} `json:"annotations"`
			Logprobs    []interface{} `json:"logprobs"`
			Text        string        `json:"text"`
		} `json:"content,omitempty"`
		Role string `json:"role,omitempty"`
	} `json:"output"`
	ParallelToolCalls  bool        `json:"parallel_tool_calls"`
	PreviousResponseID interface{} `json:"previous_response_id"`
	PromptCacheKey     interface{} `json:"prompt_cache_key"`
	Reasoning          struct {
		Effort  string      `json:"effort"`
		Summary interface{} `json:"summary"`
	} `json:"reasoning"`
	SafetyIdentifier interface{} `json:"safety_identifier"`
	ServiceTier      string      `json:"service_tier"`
	Store            bool        `json:"store"`
	Temperature      float64     `json:"temperature"`
	Text             struct {
		Format struct {
			Type string `json:"type"`
		} `json:"format"`
		Verbosity string `json:"verbosity"`
	} `json:"text"`
	ToolChoice  string        `json:"tool_choice"`
	Tools       []interface{} `json:"tools"`
	TopLogprobs int           `json:"top_logprobs"`
	TopP        float64       `json:"top_p"`
	Truncation  string        `json:"truncation"`
	Usage       struct {
		InputTokens        int `json:"input_tokens"`
		InputTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
		} `json:"input_tokens_details"`
		OutputTokens        int `json:"output_tokens"`
		OutputTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"output_tokens_details"`
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
	User     interface{} `json:"user"`
	Metadata struct {
	} `json:"metadata"`
}

func (r *ChatGPTResponse) GetText() string {
	res := ""

	for _, out := range r.Output {
		for _, content := range out.Content {
			res += content.Text
		}
	}

	return res
}
