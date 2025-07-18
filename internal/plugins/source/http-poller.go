package source

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"net/http"
	"strings"
	"time"

	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
)

type HttpPollSource struct {
	name     string
	url      string
	method   string
	headers  map[string]string
	body     string
	interval time.Duration
	client   *http.Client
	stopChan chan struct{}
}

type HttpPollConfig struct {
	URL      string            `mapstructure:"url"`
	Method   string            `mapstructure:"method"`
	Headers  map[string]string `mapstructure:"headers,omitempty"`
	Body     string            `mapstructure:"body,omitempty"`
	Interval string            `mapstructure:"interval"` // duration string like "5s", "1m"
}

func NewHttpPollSource(name string, cfg config.SourceConfig) (*HttpPollSource, error) {
	var sc HttpPollConfig
	if err := mapstructure.Decode(cfg.Config, &sc); err != nil {
		return nil, err
	}

	if sc.URL == "" || sc.Method == "" || sc.Interval == "" {
		return nil, errors.New("missing required fields in http poll config")
	}

	interval, err := time.ParseDuration(sc.Interval)
	if err != nil {
		return nil, err
	}

	return &HttpPollSource{
		name:     name,
		url:      sc.URL,
		method:   strings.ToUpper(sc.Method),
		headers:  sc.Headers,
		body:     sc.Body,
		interval: interval,
		client:   &http.Client{},
		stopChan: make(chan struct{}),
	}, nil
}

func (h *HttpPollSource) Start(ctx context.Context, eventChan chan<- config.Event) error {
	go func() {
		ticker := time.NewTicker(h.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				req, err := http.NewRequest(h.method, h.url, strings.NewReader(h.body))
				if err != nil {
					log.Printf("HttpPollSource[%s] request error: %v", h.name, err)
					continue
				}
				for k, v := range h.headers {
					req.Header.Set(k, v)
				}

				resp, err := h.client.Do(req)
				if err != nil {
					log.Printf("HttpPollSource[%s] request failed: %v", h.name, err)
					continue
				}

				bodyBytes, err := io.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					log.Printf("HttpPollSource[%s] failed reading response: %v", h.name, err)
					continue
				}

				// Try decoding as JSON
				var data map[string]interface{}
				if err := json.Unmarshal(bodyBytes, &data); err != nil {
					data = map[string]interface{}{
						"raw": string(bodyBytes),
					}
				}

				event := config.Event{
					ID:      uuid.New().String(),
					Source:  h.name,
					Data:    data,
					Headers: map[string]string{"StatusCode": resp.Status},
				}

				

				select {
				case eventChan <- event:
				case <-ctx.Done():
					return
				}

			case <-ctx.Done():
				return
			case <-h.stopChan:
				return
			}
		}
	}()

	return nil
}

func (h *HttpPollSource) Stop() error {
	close(h.stopChan)
	return nil
}

func (h *HttpPollSource) GetName() string {
	return h.name
}
