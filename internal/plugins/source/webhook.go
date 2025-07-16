package source

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)


type WebhookSource struct {
	name   string
	path   string
	method string
	server *http.Server
}


type WebhookSourceConfig struct {
	Path   string `mapstructure:"path"`
	Method string `mapstructure:"method"`
	Schema string `mapstructure:"schema,omitempty"`
}

func NewWebhookSource(name string, cfg config.SourceConfig) (*WebhookSource, error) {
	var sc WebhookSourceConfig
	if err := mapstructure.Decode(cfg.Config, &sc); err != nil {
		return nil,err
	}
	
	return &WebhookSource{
		name:   name,
		path:   sc.Path,
		method: strings.ToUpper(sc.Method),
	},nil
}

func (w *WebhookSource) Start(ctx context.Context, eventChan chan<- config.Event) error {
	mux := http.NewServeMux()
	
	mux.HandleFunc(w.path, func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != w.method {
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse request body
		var data map[string]interface{}
		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
			http.Error(rw, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Extract headers
		headers := make(map[string]string)
		for k, v := range req.Header {
			if len(v) > 0 {
				headers[k] = v[0]
			}
		}

		// Create event
		event := config.Event{
			ID:      uuid.New().String(),
			Source:  w.name,
			Data:    data,
			Headers: headers,
		}

		// Send to event channel
		select {
		case eventChan <- event:
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("OK"))
		case <-ctx.Done():
			http.Error(rw, "Server shutting down", http.StatusServiceUnavailable)
		}
	})

	w.server = &http.Server{
		Addr:    ":8080", // You can make this configurable
		Handler: mux,
	}

	go func() {
		if err := w.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Webhook source %s failed: %v", w.name, err)
		}
	}()

	return nil
}

func (w *WebhookSource) Stop() error {
	if w.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return w.server.Shutdown(ctx)
	}
	return nil
}

func (w *WebhookSource) GetName() string {
	return w.name
}
