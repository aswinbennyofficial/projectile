package orchestrator

import (
	"context"
	"fmt"

	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/aswinbennyofficial/projectile/internal/plugins"
)


// Controller orchestrates the flow of events from sources to sinks.
type Controller struct {
	registry   *plugins.Registry
	
	eventChan  chan config.Event
	routes     []config.RouteEntry
}

// NewController creates and returns a new Controller instance with a buffered event channel.
func NewController(eventChanSize int) *Controller {
	return &Controller{
		registry:  plugins.NewRegistry(),
		eventChan: make(chan config.Event, eventChanSize),
	}
}

// Initialize sets up sources and sinks using provided infrastructure config.
func (c *Controller) Initialize(infraConfig *config.InfraConfig, routesConfig *config.RoutesConfig) error {
	
	// Initialize sources and sinks
	if err := c.registry.InitializeSources(infraConfig.Sources); err != nil {
		return fmt.Errorf("failed to initialize sources: %w", err)
	}
	if err := c.registry.InitializeSinks(infraConfig.Sinks); err != nil {
		return fmt.Errorf("failed to initialize sinks: %w", err)
	}

	c.routes = routesConfig.Routes
	return nil
}


// Start runs all source plugins and begins processing events.
func (c *Controller) Start(ctx context.Context) error {
	// Start all sources
	for _, source := range c.registry.GetAllSources() {
		if err := source.Start(ctx, c.eventChan); err != nil {
			return fmt.Errorf("failed to start source %s: %w", source.GetName(), err)
		}
	}

	// Start event processor
	go c.processEvents(ctx)

	return nil
}



