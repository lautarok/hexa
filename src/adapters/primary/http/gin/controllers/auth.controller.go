package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/middlewares"
	"github.com/lautarok/hexa/src/application/domain"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
	"github.com/lautarok/hexa/src/application/usecases"
)

type AuthController struct {
	authMiddleware *middlewares.AuthMiddleware
	signupUsecase  *usecases.SignupUsecase
	loginUsecase   *usecases.LoginUsecase
	validation     ports.ValidationPort
}

type AuthControllerDeps struct {
	AuthMiddleware *middlewares.AuthMiddleware
	SignupUsecase  *usecases.SignupUsecase
	LoginUsecase   *usecases.LoginUsecase
	Validation     ports.ValidationPort
}

func NewAuthController(deps *AuthControllerDeps) *AuthController {
	return &AuthController{
		authMiddleware: deps.AuthMiddleware,
		signupUsecase:  deps.SignupUsecase,
		loginUsecase:   deps.LoginUsecase,
		validation:     deps.Validation,
	}
}

func (controller *AuthController) Name() string {
	return "auth"
}

func (controller *AuthController) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.POST("signup", controller.Signup)
	group.POST("login", controller.Login)
	group.GET("", controller.authMiddleware.HandleAuth(), controller.GetMyUser)
}

func (controller *AuthController) Signup(ctx *gin.Context) {
	var reqBody dtos.SignupInputDto
	ctx.ShouldBindBodyWithJSON(&reqBody)

	if err := reqBody.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	usecaseInput := &usecases.SignupUsecaseInput{
		Name:     reqBody.Name,
		Surname:  reqBody.Surname,
		Email:    reqBody.Email,
		Username: reqBody.Username,
		Password: reqBody.Password,
	}

	token, appErr := controller.signupUsecase.Signup(ctx, usecaseInput)
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(http.StatusCreated, &dtos.TokenOutputDto{
		Token: token.Token,
		Exp:   token.Exp,
	})
}

func (controller *AuthController) Login(ctx *gin.Context) {
	var reqBody dtos.LoginInputDto
	ctx.ShouldBindBodyWithJSON(&reqBody)

	if err := reqBody.Validate(controller.validation); err != nil {
		ctx.Error(errors.NewInvalidInputError(err.Error()))
		return
	}

	usecaseInput := &usecases.LoginUsecaseInput{
		UsernameOrEmail: reqBody.UsernameOrEmail,
		Password:        reqBody.Password,
	}

	token, err := controller.loginUsecase.Login(ctx, usecaseInput)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, &dtos.TokenOutputDto{
		Token: token.Token,
		Exp:   token.Exp,
	})
}

func (controller *AuthController) GetMyUser(ctx *gin.Context) {
	userCtx, exists := ctx.Get("auth_user")

	if !exists {
		ctx.Error(
			errors.NewNotFoundError("User not found"),
		)
		return
	}

	user, ok := userCtx.(*domain.User)
	if !ok {
		ctx.Error(
			errors.NewNotFoundError("User not found"),
		)
		return
	}

	permissions := []string{}
	for _, permission := range user.Role.Permissions {
		permissions = append(permissions, permission.Alias)
	}

	ctx.JSON(http.StatusOK, &dtos.UserOutputDto{
		ID:        user.ID,
		Name:      user.Name,
		Surname:   user.Surname,
		CreatedAt: user.CreatedAt,
		Credential: &dtos.CredentialOutputDto{
			ID:          user.Credential.ID,
			Username:    user.Credential.Username,
			Email:       user.Credential.Email,
			Permissions: permissions,
		},
	})
}
