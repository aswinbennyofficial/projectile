package plugins

import (
	"fmt"
	"sync"

	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/aswinbennyofficial/projectile/internal/plugins/sink"
	"github.com/aswinbennyofficial/projectile/internal/plugins/source"
)

type Registry struct {
	sources map[string]source.Source
	sinks   map[string]sink.Sink
	mu      sync.RWMutex
}

func NewRegistry() *Registry {
	return &Registry{
		sources: make(map[string]source.Source),
		sinks:   make(map[string]sink.Sink),
	}
}

// InitializeSources creates all sources from config
func (r *Registry) InitializeSources(sources map[string]config.SourceConfig) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for name, cfg := range sources {
		source, err := r.createSource(name, cfg)
		if err != nil {
			return fmt.Errorf("failed to create source %s: %w", name, err)
		}
		r.sources[name] = source
	}
	return nil
}

// InitializeSinks creates all sinks from config
func (r *Registry) InitializeSinks(sinks map[string]config.SinkConfig) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for name, cfg := range sinks {
		sink, err := r.createSink(name, cfg)
		if err != nil {
			return fmt.Errorf("failed to create sink %s: %w", name, err)
		}
		r.sinks[name] = sink
	}
	return nil
}



func (r *Registry) createSource(name string, cfg config.SourceConfig) (source.Source, error) {
	constructor, ok := SourceFactories[cfg.Type]
	if !ok {
		return nil, fmt.Errorf("unknown source type: %s", cfg.Type)
	}
	return constructor(name, cfg)
}

func (r *Registry) createSink(name string, cfg config.SinkConfig) (sink.Sink, error) {
	constructor, ok := SinkFactories[cfg.Type]
	if !ok {
		return nil, fmt.Errorf("unknown sink type: %s", cfg.Type)
	}
	return constructor(name, cfg)
}



// GetSource returns a source by name
func (r *Registry) GetSource(name string) (source.Source, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	source, exists := r.sources[name]
	return source, exists
}

// GetSink returns a sink by name
func (r *Registry) GetSink(name string) (sink.Sink, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	sink, exists := r.sinks[name]
	return sink, exists
}

// GetAllSources returns all sources
func (r *Registry) GetAllSources() map[string]source.Source {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make(map[string]source.Source)
	for k, v := range r.sources {
		result[k] = v
	}
	return result
}