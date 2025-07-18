package sink

import (
	"context"

	"github.com/aswinbennyofficial/projectile/internal/config"
)


// Sink is an interface that represents an output target (e.g., database, HTTP endpoint, file, etc.)
// Any plugin that wants to act as a sink must implement this interface.
type Sink interface {
	// Send handles writing or delivering the event to the target system.
	Send(ctx context.Context, event config.Event) error
	
	// GetName returns the unique name of sink
	GetName() string
}