package plugins

import (
	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/aswinbennyofficial/projectile/internal/plugins/sink"
	"github.com/aswinbennyofficial/projectile/internal/plugins/source"
)



/*
MAKE CHANGES HERE TO ADD OR REMOVE NEW PLUGINS
*/


// SinkFactories is a registry that maps sink types (e.g., "stdout", "file", "webhook")
// to their respective constructor functions. These functions take the name and
// configuration of the sink and return an initialized Sink instance.
var SinkFactories = map[string]func(name string, cfg config.SinkConfig) (sink.Sink, error){
	"stdout": func(name string, cfg config.SinkConfig) (sink.Sink, error) {
		return sink.NewStdoutSink(name, cfg)
	},
	"file": func(name string, cfg config.SinkConfig) (sink.Sink, error) {
		return sink.NewFileSink(name, cfg)
	},
	"http": func(name string, cfg config.SinkConfig) (sink.Sink, error) {
		return sink.NewHttpSink(name, cfg)
	},
}




// SourceFactories is a registry that maps source types (e.g., "webhook")
// to their respective constructor functions. These functions take the name and
// configuration of the source and return an initialized Source instance.
var SourceFactories = map[string]func(name string, cfg config.SourceConfig) (source.Source, error){
	"webhook": func(name string, cfg config.SourceConfig) (source.Source, error) {
		return source.NewWebhookSource(name, cfg)
	},
}
