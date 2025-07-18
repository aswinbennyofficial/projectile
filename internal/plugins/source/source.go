package source

import (
	"context"

	"github.com/aswinbennyofficial/projectile/internal/config"
)

// Source is an interface that represents an input event generator (e.g., webhook, Kafka, file watcher, etc.)
// Any plugin that wants to act as a source must implement this interface.
type Source interface {
	// Start begins the source's event listening or generation and pushes events into eventChan.
	Start(ctx context.Context, eventChan chan<- config.Event) error
	
	// Stop gracefully shuts down the source, cleaning up any resources (if needed).
	Stop() error

	// GetName returns the unique name or identifier of the source.
	GetName() string
}