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
	Type    string            `mapstructure:"type"`
	Path    string            `mapstructure:"path,omitempty"`
	Method  string            `mapstructure:"method,omitempty"`
	Schema  string            `mapstructure:"schema,omitempty"`
	DSN     string            `mapstructure:"dsn,omitempty"`
	Headers map[string]string `mapstructure:"headers,omitempty"`
}

type SinkConfig struct {
	Type    string            `mapstructure:"type"`
	Path    string            `mapstructure:"path,omitempty"`
	Method  string            `mapstructure:"method,omitempty"`
	URL     string            `mapstructure:"url,omitempty"`
	DSN     string            `mapstructure:"dsn,omitempty"`
	Topic   string            `mapstructure:"topic,omitempty"`
	Headers map[string]string `mapstructure:"headers,omitempty"`
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