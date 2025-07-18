package orchestrator

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/aswinbennyofficial/projectile/internal/config"
	
)

// routeEvent routes a single event to all matching sinks based on the route config.
func (c *Controller) routeEvent(ctx context.Context, event config.Event) {
	for _, route := range c.routes {
		// Check if the route is applicable for the event's source
		if route.Source == event.Source {
			for _, sinkName := range route.Sinks {
				// Retrieve the sink and send the event
				if sink, exists := c.registry.GetSink(sinkName); exists {
					if err := sink.Send(ctx, event); err != nil {
						log.Printf("Failed to send event to sink %s: %v", sinkName, err)
					}
				}
			}
		}
	}
}
