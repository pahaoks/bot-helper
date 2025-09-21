package repositories

import (
	"bot-helper/internal/domain/entities"
	"bot-helper/pkg/logger"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrUnauthorized = errors.New("unauthorized")

type AnkiWebConfig struct {
	BaseURL string `default:"http://localhost:8765"`
}

type AnkiWebRepository struct {
	config     AnkiWebConfig
	httpClient *http.Client
	logger     logger.Logger
}

func NewAnkiWebRepository(
	config AnkiWebConfig,
	rt http.RoundTripper,
	logger logger.Logger,
) *AnkiWebRepository {
	return &AnkiWebRepository{
		config:     config,
		httpClient: &http.Client{Transport: rt},
		logger:     logger,
	}
}

func (r *AnkiWebRepository) AddNote(
	deckName, modelName string, front string, back string,
) error {
	req := entities.AnkiRequest{
		Action:  "addNote",
		Version: 6,
	}
	req.Params.Note.DeckName = deckName
	req.Params.Note.ModelName = modelName
	req.Params.Note.Fields.Front = front
	req.Params.Note.Fields.Back = back
	req.Params.Note.Options.AllowDuplicate = false
	req.Params.Note.Options.DuplicateScope = "deck"
	req.Params.Note.Options.DuplicateScopeOptions.DeckName = deckName
	req.Params.Note.Options.DuplicateScopeOptions.CheckChildren = false
	req.Params.Note.Options.DuplicateScopeOptions.CheckAllModels = false
	req.Params.Note.Tags = []string{"generated_by_anki_memo_app"}

	// Send the request
	res, err := r.basePost(
		"/",
		req,
	)

	if err != nil {
		return err
	}

	if res.StatusCode >= 399 {
		return ErrUnauthorized
	}

	m := make(map[string]any)
	b, err := json.Marshal(res.Body)
	if err != nil {
		return err
	}
	r.logger.Info("AnkiWeb response: %s", string(b))
	err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&m)
	if err != nil {
		return err
	}

	if m["error"] != nil {
		return fmt.Errorf("anki error: %v", m["error"])
	}

	return nil
}

func (r *AnkiWebRepository) Sync() error {
	req := map[string]any{
		"action":  "sync",
		"version": 6,
	}

	// Send the request
	res, err := r.basePost(
		"/",
		req,
	)

	if err != nil {
		return err
	}

	if res.StatusCode >= 399 {
		return ErrUnauthorized
	}

	m := make(map[string]any)
	b, err := json.Marshal(res.Body)
	if err != nil {
		return err
	}
	r.logger.Info("AnkiWeb response: %s", string(b))
	err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&m)
	if err != nil {
		return err
	}

	if m["error"] != nil {
		return fmt.Errorf("anki error: %v", m["error"])
	}

	return nil
}

func (r *AnkiWebRepository) basePost(
	path string,
	body any,
) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	r.logger.Info("AnkiWeb request: %s, body: %s", r.config.BaseURL+path, string(b))

	req, err := http.NewRequest(
		http.MethodPost,
		r.config.BaseURL+path,
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, err
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp, nil
}
