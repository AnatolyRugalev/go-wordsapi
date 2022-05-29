package wordsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WordsAPI struct {
	options
	apiKey string
}

func New(apiKey string, opts ...Option) *WordsAPI {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}
	return &WordsAPI{
		options: o,
		apiKey:  apiKey,
	}
}

func (w *WordsAPI) WordEndpointURL(word, endpoint string) string {
	url := w.baseUrl + "/words/" + word
	if endpoint != "" {
		url += "/" + endpoint
	}
	return url
}

func (w *WordsAPI) Word(ctx context.Context, word string) (*WordResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, w.WordEndpointURL(word, ""), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating http.Request: %w", err)
	}
	req.Header.Set("X-RapidAPI-Host", "wordsapiv1.p.rapidapi.com")
	req.Header.Set("X-RapidAPI-Key", w.apiKey)
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting word: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// TODO: proper error handling: unauthorized, word not found
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	response := WordResponse{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}
	return &response, nil
}
