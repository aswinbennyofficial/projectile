package config

// InfraConfig defines the global infrastructure configuration.
type InfraConfig struct {
	Version string                    `mapstructure:"version"`
	Sources map[string]SourceConfig  `mapstructure:"sources"`
	Sinks   map[string]SinkConfig    `mapstructure:"sinks"`
}

// RoutesConfig defines routing rules for events.
type RoutesConfig struct {
	Version string       `mapstructure:"version"`
	Routes  []RouteEntry `mapstructure:"routes"`
}

// SourceConfig represents the configuration for a single source plugin.
type SourceConfig struct {
	Type   string                 `mapstructure:"type"`
	Config map[string]interface{} `mapstructure:"config"` // plugin-specific config
}

// SinkConfig represents the configuration for a single sink plugin.
type SinkConfig struct {
	Type   string                 `mapstructure:"type"`
	Config map[string]interface{} `mapstructure:"config"` // plugin-specific config
}


// RouteEntry defines how to connect one source to multiple sinks.
type RouteEntry struct {
	Name   string   `mapstructure:"name"`
	Source string   `mapstructure:"source"`
	Rules  []Rule `mapstructure:"rules"`
}

type Rule struct {
	Condition string `mapstructure:"condition"`
	Sinks     []string `mapstructure:"sinks"`
}




// Event represents data flowing through the system
type Event struct {
	ID      string                 `json:"id"`
	Source  string                 `json:"source"`
	Data    map[string]interface{} `json:"data"`
	Headers map[string]string      `json:"headers"`
}