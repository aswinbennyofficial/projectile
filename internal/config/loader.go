package config

import (
	"fmt"
	"github.com/spf13/viper"
)


// LoadInfraConfig loads and returns the infrastructure config (one-time)
func LoadInfraConfig(path string) (*InfraConfig, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("reading infra config: %w", err)
	}

	var cfg InfraConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal infra config: %w", err)
	}

	if cfg.Version != "v1" {
		return nil, fmt.Errorf("unsupported infra config version: %s", cfg.Version)
	}

	return &cfg, nil
}

// LoadRoutesConfig loads and returns the routing config (one-time)
func LoadRoutesConfig(path string) (*RoutesConfig, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("reading routes config: %w", err)
	}

	var cfg RoutesConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal routes config: %w", err)
	}

	if cfg.Version != "v1" {
		return nil, fmt.Errorf("unsupported routes config version: %s", cfg.Version)
	}

	return &cfg, nil
}
