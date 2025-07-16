package sink

import (
	"context"

	"github.com/aswinbennyofficial/projectile/internal/config"
)



type Sink interface {
	Send(ctx context.Context, event config.Event) error
	GetName() string
}