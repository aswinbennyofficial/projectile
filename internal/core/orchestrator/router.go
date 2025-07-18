package orchestrator

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/expr-lang/expr"
	
)


// routeEvent routes a single event to all matching sinks based on the route config.
func (c *Controller) routeEvent(ctx context.Context, event config.Event) {
	for _, route := range c.routes {
		if route.Source != event.Source {
			continue
		}

// routeEvent routes a single event to all matching sinks based on the route config.
		for _, rule := range route.Rules {
			match := true
			if rule.Condition != "" {
				env := map[string]interface{}{
					"event": event.Data,
				}
				result, err := expr.Eval(rule.Condition, env)
				if err != nil || result != true {
					match = false
				}
			}

			log.Info().Msgf("%s | %v",rule.Condition,match)

			if match {
				for _, sinkName := range rule.Sinks {
					if sink, exists := c.registry.GetSink(sinkName); exists {
						if err := sink.Send(ctx, event); err != nil {
							log.Printf("Failed to send event to sink %s: %v", sinkName, err)
						}
					}
				}
			}
		}
	}
}
