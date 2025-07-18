package sink

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/mitchellh/mapstructure"
)



type FileSink struct {
	name string
	path string
	append bool
}


type FileSinkConfig struct {
	Path string `mapstructure:"path"`
	Append bool   `mapstructure:"append"`
}


func NewFileSink(name string, cfg config.SinkConfig) (*FileSink,error) {
	var fc FileSinkConfig
	if err := mapstructure.Decode(cfg.Config, &fc); err != nil {
		return nil, err
	}


	return &FileSink{
		name: name,
		path: fc.Path,
		append: fc.Append,
	},nil
}

func (f *FileSink) Send(ctx context.Context, event config.Event) error {
	// Ensure directory exists
	if err := os.MkdirAll(f.path, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Marshal event to JSON
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	if f.append {
		// Append to single file: events.log
		logFilePath := filepath.Join(f.path, "events.log")
		fh, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
		defer fh.Close()

		if _, err := fh.Write(append(data, '\n')); err != nil {
			return fmt.Errorf("failed to write log entry: %w", err)
		}
	} else {
		// Create a separate file for each event
		filename := fmt.Sprintf("%s_%s.json",
			time.Now().Format("20060102_150405"),
			event.ID,
		)
		filePath := filepath.Join(f.path, filename)

		if err := os.WriteFile(filePath, data, 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}

	return nil
}


func (f *FileSink) GetName() string {
	return f.name
}
