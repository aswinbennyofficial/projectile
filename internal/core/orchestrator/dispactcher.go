package orchestrator

import "context"

// processEvents continuously listens on the event channel and dispatches events.
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