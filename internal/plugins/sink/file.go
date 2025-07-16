package sink

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aswinbennyofficial/projectile/internal/config"

)



type FileSink struct {
	name string
	path string
}

func NewFileSink(name string, cfg config.SinkConfig) *FileSink {
	return &FileSink{
		name: name,
		path: cfg.Path,
	}
}

func (f *FileSink) Send(ctx context.Context, event config.Event) error {
	// Ensure directory exists
	if err := os.MkdirAll(f.path, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create filename with timestamp
	filename := fmt.Sprintf("%s_%s.json", 
		time.Now().Format("20060102_150405"), 
		event.ID)
	
	filePath := filepath.Join(f.path, filename)
	
	// Marshal event to JSON
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Write to file
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (f *FileSink) GetName() string {
	return f.name
}
