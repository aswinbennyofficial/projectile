package source

import (
	"context"

	"github.com/aswinbennyofficial/projectile/internal/config"
)


type Source interface {
	Start(ctx context.Context, eventChan chan<- config.Event) error
	Stop() error
	GetName() string
}