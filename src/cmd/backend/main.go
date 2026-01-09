package main

import (
	"log"

	"github.com/lautarok/hexa/src/adapters/config"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/controllers"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/middlewares"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/repositories"
	"github.com/lautarok/hexa/src/adapters/secondary/validation/validator"
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

	httpAdapter := gin.NewGinAdapter("api/v1", environment == "PROD")

	errorMiddleware := middlewares.NewErrorMiddleware()
	httpAdapter.RegisterGlobalMiddlewares(errorMiddleware.HandleErrors)

	healthController := controllers.NewHealthController()

	dsn, err := envAdapter.GetStr("POSTGRES_DSN")
	if err != nil {
		log.Fatalf("Error getting POSTGRES_DSN environment variable: %v", err)
	}

	persistenceAdapter := bun.NewBunAdapter(&bun.BunAdapterDeps{
		DSN: dsn,
	})

	validation := validator.NewValidatorAdapter()

	usersRepository := repositories.NewUsersRepository(&repositories.UsersRepositoryDeps{
		DBAdapter: persistenceAdapter,
	})
	getUsersUsecase := usecases.NewGetUsersUsecase(&usecases.GetUsersUsecaseDeps{
		UsersRepository: usersRepository,
	})
	usersController := controllers.NewUsersController(&controllers.UsersControllerDeps{
		GetUsersUsecase: getUsersUsecase,
	})

	credentialsRepository := repositories.NewCredentialsRepository(&repositories.CredentialsRepositoryDeps{
		DBAdapter: persistenceAdapter,
	})
	signupUsecase := usecases.NewSignupUsecase(&usecases.SignupUsecaseDeps{
		CredentialsRepository: credentialsRepository,
		UsersRepository:       usersRepository,
		PersistenceAdapter:    persistenceAdapter,
	})
	authController := controllers.NewAuthController(&controllers.AuthControllerDeps{
		SignupUsecase: signupUsecase,
		Validation:    validation,
	})

	httpAdapter.RegisterControllers(
		healthController,
		usersController,
		authController,
	)

	httpPort, err := envAdapter.GetStr("HTTP_PORT")
	if err != nil {
		log.Fatalf("Error getting HTTP_PORT environment variable: %v", err)
	}

	httpAdapter.Start(":" + httpPort)
}
