package repositories

import (
	"bot-helper/internal/domain/entities"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type ChatGPTConfig struct {
	BaseURL           string `default:"https://api.openai.com"`
	APIKey            string `default:""`
	Model             string `default:"gpt-5-mini"`
	ModelToTranscribe string `default:"gpt-4o-transcribe"`
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

func (r *ChatGPTRepository) TranscribeAudio(
	file *os.File,
) (entities.ChatGPTTranscriptionResponse, error) {
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)

	// Add the file field
	fileWriter, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return entities.ChatGPTTranscriptionResponse{}, err
	}

	// Copy the file content to the form file field
	if _, err := io.Copy(fileWriter, file); err != nil {
		return entities.ChatGPTTranscriptionResponse{}, err
	}

	// Add the model field
	if err := writer.WriteField("model", r.config.ModelToTranscribe); err != nil {
		return entities.ChatGPTTranscriptionResponse{}, err
	}

	// Close the multipart writer to set the terminating boundary
	if err := writer.Close(); err != nil {
		return entities.ChatGPTTranscriptionResponse{}, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST",
		r.config.BaseURL+"/v1/audio/transcriptions",
		reqBody,
	)
	if err != nil {
		return entities.ChatGPTTranscriptionResponse{}, err
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.config.APIKey))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return entities.ChatGPTTranscriptionResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.ChatGPTTranscriptionResponse{}, err
	}

	res := entities.ChatGPTTranscriptionResponse{}
	err = json.Unmarshal(respBody, &res)

	return res, err
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
