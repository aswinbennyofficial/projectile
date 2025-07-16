package plugins

import (
	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/aswinbennyofficial/projectile/internal/plugins/sink"
	"github.com/aswinbennyofficial/projectile/internal/plugins/source"
)



// SinkFactories maps sink type → constructor
var SinkFactories = map[string]func(name string, cfg config.SinkConfig) (sink.Sink, error){
	"stdout": func(name string, cfg config.SinkConfig) (sink.Sink, error) {
		return sink.NewStdoutSink(name, cfg), nil
	},
	"file": func(name string, cfg config.SinkConfig) (sink.Sink, error) {
		return sink.NewFileSink(name, cfg), nil
	},
	"webhook": func(name string, cfg config.SinkConfig) (sink.Sink, error) {
		return sink.NewWebhookSink(name, cfg), nil
	},
}



// SourceFactories maps source type → constructor
var SourceFactories = map[string]func(name string, cfg config.SourceConfig) (source.Source, error){
	"webhook": func(name string, cfg config.SourceConfig) (source.Source, error) {
		return source.NewWebhookSource(name, cfg), nil
	},
}
