package sink

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/mitchellh/mapstructure"
)




type HttpSink struct {
	name    string
	url     string
	method  string
	headers map[string]string
	client  *http.Client
}


type HttpSinkConfig struct {
	URL     string            `mapstructure:"url"`
	Method  string            `mapstructure:"method"`
	Headers map[string]string `mapstructure:"headers"`
}

func NewHttpSink(name string, cfg config.SinkConfig) (*HttpSink, error) {
	var wc HttpSinkConfig
	if err := mapstructure.Decode(cfg.Config, &wc); err != nil {
		return nil,err
	}

	return &HttpSink{
		name:    name,
		url:     wc.URL,
		method:  wc.Method,
		headers: wc.Headers,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	},nil
}

func (w *HttpSink) Send(ctx context.Context, event config.Event) error {
	// Marshal event to JSON
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, w.method, w.url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	for k, v := range w.headers {
		req.Header.Set(k, v)
	}

	// Send request
	resp, err := w.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode >= 400 {
		return fmt.Errorf("http request returned status %d", resp.StatusCode)
	}

	return nil
}

func (w *HttpSink) GetName() string {
	return w.name
}
