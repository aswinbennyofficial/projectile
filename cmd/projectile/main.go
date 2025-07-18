package main

import (
	"context"

	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/aswinbennyofficial/projectile/internal/core/orchestrator"
	"github.com/aswinbennyofficial/projectile/internal/utils"
)

func main() {
	ctx := context.Background()
	eventChanSize := 200  // Size of the buffered event channel


	// Initialize logger
	logger := utils.NewLogger()
	logger.Info().Msg("🚀 Starting Projectile...")

	// Load infrastructure configuration from YAML file
	infra, err := config.LoadInfraConfig("configs/infra.yaml")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load infra config:")
	}

	// Load routing configuration from YAML file
	routes, err := config.LoadRoutesConfig("configs/routes.yaml")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load routes config:")
	}

	// Create a new orchestrator controller with specified event channel size
	ctrl := orchestrator.NewController(eventChanSize)

	// Initialize the orchestrator controller with loaded configs
	if err := ctrl.Initialize(infra, routes); err != nil {
		logger.Fatal().Err(err).Msg("failed to initialize controller:")
	}

	// Start the orchestrator event loop
	if err := ctrl.Start(ctx); err != nil {
		logger.Fatal().Err(err).Msg("failed to start controller:")
	}

	// Block forever
	select {}
}
