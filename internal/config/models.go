package config

type InfraConfig struct {
	Version string                    `mapstructure:"version"`
	Sources map[string]SourceConfig  `mapstructure:"sources"`
	Sinks   map[string]SinkConfig    `mapstructure:"sinks"`
}

type RoutesConfig struct {
	Version string       `mapstructure:"version"`
	Routes  []RouteEntry `mapstructure:"routes"`
}

type SourceConfig struct {
	Type   string                 `mapstructure:"type"`
	Config map[string]interface{} `mapstructure:"config"` // plugin-specific config
}

type SinkConfig struct {
	Type   string                 `mapstructure:"type"`
	Config map[string]interface{} `mapstructure:"config"` // plugin-specific config
}







type RouteEntry struct {
	Name   string   `mapstructure:"name"`
	Source string   `mapstructure:"source"`
	Sinks  []string `mapstructure:"sinks"`
}


// Event represents data flowing through the system
type Event struct {
	ID      string                 `json:"id"`
	Source  string                 `json:"source"`
	Data    map[string]interface{} `json:"data"`
	Headers map[string]string      `json:"headers"`
}