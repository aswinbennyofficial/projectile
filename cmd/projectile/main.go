package main

import (
	"context"

	"github.com/aswinbennyofficial/projectile/internal/config"
	"github.com/aswinbennyofficial/projectile/internal/core/orchestrator"
	"github.com/aswinbennyofficial/projectile/internal/utils"
)

func main() {
	ctx := context.Background()

	// Initialize logger
	logger := utils.NewLogger()
	logger.Info().Msg("🚀 Starting Projectile...")

	infra, err := config.LoadInfraConfig("configs/infra.yaml")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load infra config:")
	}

	routes, err := config.LoadRoutesConfig("configs/routes.yaml")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load routes config:")
	}

	ctrl := orchestrator.NewController()
	if err := ctrl.Initialize(infra, routes); err != nil {
		logger.Fatal().Err(err).Msg("failed to initialize controller:")
	}

	if err := ctrl.Start(ctx); err != nil {
		logger.Fatal().Err(err).Msg("failed to start controller:")
	}

	// Block forever
	select {}
}
