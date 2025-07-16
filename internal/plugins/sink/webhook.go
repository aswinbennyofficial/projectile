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




type WebhookSink struct {
	name    string
	url     string
	method  string
	headers map[string]string
	client  *http.Client
}


type WebhookSinkConfig struct {
	URL     string            `mapstructure:"url"`
	Method  string            `mapstructure:"method"`
	Headers map[string]string `mapstructure:"headers"`
}

func NewWebhookSink(name string, cfg config.SinkConfig) *WebhookSink {
	var wc WebhookSinkConfig
	if err := mapstructure.Decode(cfg.Config, &wc); err != nil {
		return nil
	}

	return &WebhookSink{
		name:    name,
		url:     wc.URL,
		method:  wc.Method,
		headers: wc.Headers,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (w *WebhookSink) Send(ctx context.Context, event config.Event) error {
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
		return fmt.Errorf("webhook returned status %d", resp.StatusCode)
	}

	return nil
}

func (w *WebhookSink) GetName() string {
	return w.name
}
