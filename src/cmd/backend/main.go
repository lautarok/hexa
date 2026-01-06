package main

import (
	"log"

	"github.com/lautarok/hexa/src/adapters/config"
	"github.com/lautarok/hexa/src/adapters/primary/http"
	"github.com/lautarok/hexa/src/app/modules/health"
)

func main() {
	envAdapter := config.NewGodotEnvAdapter()
	if err := envAdapter.Load(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	healthController := health.NewHealthController(&health.HealthControllerDeps{
		HealthService: health.NewHealthService(),
	})

	environment, err := envAdapter.GetStr("ENVIRONMENT")
	if err != nil {
		environment = "DEV"
	}

	httpAdapter := http.NewGinAdapter("api/v1", environment == "PROD")
	httpAdapter.RegisterControllers(healthController)

	httpPort, err := envAdapter.GetStr("HTTP_PORT")
	if err != nil {
		log.Fatalf("Error getting HTTP_PORT environment variable: %v", err)
	}

	httpAdapter.Start(":" + httpPort)
}
