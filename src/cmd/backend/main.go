package main

import (
	"log"

	"github.com/lautarok/hexa/src/adapters/primary/http/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/controllers"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/middlewares"
	"github.com/lautarok/hexa/src/adapters/secondary/config"
	"github.com/lautarok/hexa/src/adapters/secondary/identity/jwt"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/repositories"
	"github.com/lautarok/hexa/src/adapters/secondary/validation/validator"
	authCommand "github.com/lautarok/hexa/src/application/usecases/auth/command"
	authQuery "github.com/lautarok/hexa/src/application/usecases/auth/query"
	permimssionsQuery "github.com/lautarok/hexa/src/application/usecases/permissions/query"
	rolesCommand "github.com/lautarok/hexa/src/application/usecases/roles/command"
	rolesQuery "github.com/lautarok/hexa/src/application/usecases/roles/query"
	usersQuery "github.com/lautarok/hexa/src/application/usecases/users/query"
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
		log.Fatal(err)
	}

	persistenceAdapter := bun.NewBunAdapter(&bun.BunAdapterDeps{
		DSN: dsn,
	})

	validationAdapter := validator.NewValidatorAdapter()

	jwtSecret, err := envAdapter.GetStr("JWT_SECRET")
	if err != nil {
		log.Fatal(err)
	}

	identityAdapter := jwt.NewJWTAdapter(&jwt.JWTAdapterDeps{
		Secret: jwtSecret,
	})

	usersRepository := repositories.NewUsersRepository(&repositories.UsersRepositoryDeps{
		DBAdapter: persistenceAdapter,
	})
	credentialsRepository := repositories.NewCredentialsRepository(&repositories.CredentialsRepositoryDeps{
		DBAdapter: persistenceAdapter,
	})
	rolesRepository := repositories.NewRolesRepository(&repositories.RolesRepositoryDeps{
		DBAdapter: persistenceAdapter,
	})
	permissionsRepository := repositories.NewPermissionsRepository(&repositories.PermissionsRepositoryDeps{
		DBAdapter: persistenceAdapter,
	})

	getUsersUsecase := usersQuery.NewGetUsersUsecase(&usersQuery.GetUsersUsecaseDeps{
		UsersRepository: usersRepository,
	})
	signupUsecase := authCommand.NewSignupUsecase(&authCommand.SignupUsecaseDeps{
		CredentialsRepository: credentialsRepository,
		UsersRepository:       usersRepository,
		RolesRepository:       rolesRepository,
		PersistenceAdapter:    persistenceAdapter,
		IdentityAdapter:       identityAdapter,
	})
	loginUsecase := authCommand.NewLoginUsecase(&authCommand.LoginUsecaseDeps{
		CredentialsRepository: credentialsRepository,
		IdentityAdapter:       identityAdapter,
	})
	getUserFromTokenUsecase := authQuery.NewGetUserFromTokenUsecase(&authQuery.GetUserFromTokenUsecaseDeps{
		UsersRepository:       usersRepository,
		CredentialsRepository: credentialsRepository,
		IdentityAdapter:       identityAdapter,
	})
	getRolesUsecase := rolesQuery.NewGetRolesUsecase(&rolesQuery.GetRolesUsecaseDeps{
		RolesRepository: rolesRepository,
	})
	createRolesUsecase := rolesCommand.NewCreateRoleUsecase(&rolesCommand.CreateRoleUsecaseDeps{
		RolesRepository:       rolesRepository,
		PermissionsRepository: permissionsRepository,
		PersistenceAdapter:    persistenceAdapter,
	})
	getPermissionsUsecase := permimssionsQuery.NewGetPermissionsUsecase(&permimssionsQuery.GetPermissionsUsecaseDeps{
		PermissionsRepository: permissionsRepository,
	})

	authMiddleware := middlewares.NewAuthMiddleware(&middlewares.AuthMiddlewareDeps{
		GetUserFromTokenUsecase: getUserFromTokenUsecase,
	})

	usersController := controllers.NewUsersController(&controllers.UsersControllerDeps{
		GetUsersUsecase: getUsersUsecase,
		Validation:      validationAdapter,
	})
	authController := controllers.NewAuthController(&controllers.AuthControllerDeps{
		SignupUsecase:  signupUsecase,
		LoginUsecase:   loginUsecase,
		Validation:     validationAdapter,
		AuthMiddleware: authMiddleware,
	})
	rolesController := controllers.NewRolesController(&controllers.RolesControllerDeps{
		GetRolesUsecase:   getRolesUsecase,
		CreateRoleUsecase: createRolesUsecase,
		AuthMiddleware:    authMiddleware,
		Validation:        validationAdapter,
	})
	permissionsController := controllers.NewPermissionsController(&controllers.PermissionsControllerDeps{
		GetPermissionsUsecase: getPermissionsUsecase,
		Validation:            validationAdapter,
	})

	httpAdapter.RegisterControllers(
		healthController,
		usersController,
		authController,
		rolesController,
		permissionsController,
	)

	httpPort, err := envAdapter.GetStr("HTTP_PORT")
	if err != nil {
		log.Fatal(err)
	}

	httpAdapter.Start(":" + httpPort)
}
