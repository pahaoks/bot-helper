package repositories

import (
	"bot-helper/internal/domain/entities"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ChatGPTConfig struct {
	BaseURL string `default:"https://api.openai.com"`
	APIKey  string `default:""`
	Model   string `default:"gpt-5-mini"`
}

type ChatGPTRepository struct {
	config     ChatGPTConfig
	httpClient *http.Client
}

func NewChatGPTRepository(
	config ChatGPTConfig,
	rt http.RoundTripper,
) *ChatGPTRepository {
	client := &http.Client{
		Transport: rt,
	}

	return &ChatGPTRepository{
		config:     config,
		httpClient: client,
	}
}

func (r *ChatGPTRepository) Prompt(
	prompt string,
) (entities.ChatGPTResponse, error) {
	req := entities.ChatGPTRequest{
		Model: r.config.Model,
		Input: prompt,
		Store: false,
	}

	resp, err := r.post("/v1/responses", req)
	if err != nil {
		return entities.ChatGPTResponse{}, err
	}

	return r.handleResponse(resp)
}

func (r *ChatGPTRepository) post(
	path string,
	body any,
) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		r.config.BaseURL+path,
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.config.APIKey)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ChatGPTRepository) handleResponse(
	resp *http.Response,
) (entities.ChatGPTResponse, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.ChatGPTResponse{}, err
	}

	var response entities.ChatGPTResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return entities.ChatGPTResponse{}, err
	}

	return response, nil
}
