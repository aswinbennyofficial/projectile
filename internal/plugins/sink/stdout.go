package sink

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aswinbennyofficial/projectile/internal/config"
)



type StdoutSink struct {
	name string
}

func NewStdoutSink(name string, cfg config.SinkConfig) *StdoutSink {
	return &StdoutSink{name: name}
}

func (s *StdoutSink) Send(ctx context.Context, event config.Event) error {
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	
	fmt.Printf("[%s] %s\n", s.name, string(data))
	return nil
}

func (s *StdoutSink) GetName() string {
	return s.name
}