package main

import (
	"log"

	"github.com/lautarok/hexa/src/adapters/config"
	"github.com/lautarok/hexa/src/adapters/primary/http"
	"github.com/lautarok/hexa/src/adapters/primary/http/controllers"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/repositories"
	"github.com/lautarok/hexa/src/app/usecases"
)

func main() {
	envAdapter := config.NewGodotEnvAdapter()
	if err := envAdapter.Load(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	environment, err := envAdapter.GetStr("ENVIRONMENT")
	if err != nil {
		environment = "DEV"
	}

	httpAdapter := http.NewGinAdapter("api/v1", environment == "PROD")

	healthController := controllers.NewHealthController()

	dsn, err := envAdapter.GetStr("POSTGRES_DSN")
	if err != nil {
		log.Fatalf("Error getting POSTGRES_DSN environment variable: %v", err)
	}

	dbAdapter := persistence.NewBunAdapter(&persistence.BunAdapterDeps{
		DSN: dsn,
	})
	db := dbAdapter.GetDB()

	usersRepository := repositories.NewUsersRepository(&repositories.UsersRepositoryDeps{
		DB: db,
	})
	getUsersUsecase := usecases.NewGetUsersUsecase(&usecases.GetUsersUsecaseDeps{
		UsersRepository: usersRepository,
	})
	usersController := controllers.NewUsersController(&controllers.UsersControllerDeps{
		GetUsersUsecase: getUsersUsecase,
	})

	httpAdapter.RegisterControllers(
		healthController,
		usersController,
	)

	httpPort, err := envAdapter.GetStr("HTTP_PORT")
	if err != nil {
		log.Fatalf("Error getting HTTP_PORT environment variable: %v", err)
	}

	httpAdapter.Start(":" + httpPort)
}
