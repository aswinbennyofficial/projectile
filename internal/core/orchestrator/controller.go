package orchestrator

import (
	"context"
	"fmt"
	"log"

	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/aswinbennyofficial/projectile/internal/plugins"
)

type Controller struct {
	registry   *plugins.Registry
	
	eventChan  chan config.Event
	routes     []config.RouteEntry
}

func NewController() *Controller {
	return &Controller{
		registry:  plugins.NewRegistry(),
		eventChan: make(chan config.Event, 100),
	}
}

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

func (c *Controller) processEvents(ctx context.Context) {
	for {
		select {
		case event := <-c.eventChan:
			c.routeEvent(ctx, event)
		case <-ctx.Done():
			return
		}
	}
}

func (c *Controller) routeEvent(ctx context.Context, event config.Event) {
	// Find routes for this event's source
	for _, route := range c.routes {
		if route.Source == event.Source {
			// Send to all sinks in this route
			for _, sinkName := range route.Sinks {
				if sink, exists := c.registry.GetSink(sinkName); exists {
					if err := sink.Send(ctx, event); err != nil {
						log.Printf("Failed to send event to sink %s: %v", sinkName, err)
					}
				}
			}
		}
	}
}