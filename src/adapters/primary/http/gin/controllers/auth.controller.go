package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/middlewares"
	authCommand "github.com/lautarok/hexa/src/application/usecases/auth/command"
	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/lautarok/hexa/src/domain/ports"
)

type AuthController struct {
	authMiddleware *middlewares.AuthMiddleware
	signupUsecase  *authCommand.SignupUsecase
	loginUsecase   *authCommand.LoginUsecase
	validation     ports.ValidationPort
}

type AuthControllerDeps struct {
	AuthMiddleware *middlewares.AuthMiddleware
	SignupUsecase  *authCommand.SignupUsecase
	LoginUsecase   *authCommand.LoginUsecase
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
	ctx.ShouldBindJSON(&reqBody)

	if err := reqBody.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	usecaseInput := &authCommand.SignupUsecaseInput{
		Name:     reqBody.Name,
		Surname:  reqBody.Surname,
		Email:    reqBody.Email,
		Username: reqBody.Username,
		Password: reqBody.Password,
	}

	result, appErr := controller.signupUsecase.Signup(ctx, usecaseInput)
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		dtos.NewTokenOutputDto(result.Token, result.Exp, result.User),
	)
}

func (controller *AuthController) Login(ctx *gin.Context) {
	var reqBody dtos.LoginInputDto
	ctx.ShouldBindJSON(&reqBody)

	if err := reqBody.Validate(controller.validation); err != nil {
		ctx.Error(errors.NewInvalidInputError(err.Error()))
		return
	}

	usecaseInput := &authCommand.LoginUsecaseInput{
		UsernameOrEmail: reqBody.UsernameOrEmail,
		Password:        reqBody.Password,
	}

	result, err := controller.loginUsecase.Login(ctx, usecaseInput)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		dtos.NewTokenOutputDto(result.Token, result.Exp, result.User),
	)
}

func (controller *AuthController) GetMyUser(ctx *gin.Context) {
	userCtx, exists := ctx.Get("auth_user")

	if !exists {
		ctx.Error(
			errors.NewNotFoundError("User not found"),
		)
		return
	}

	user, ok := userCtx.(*models.User)
	if !ok {
		ctx.Error(
			errors.NewNotFoundError("User not found"),
		)
		return
	}

	ctx.JSON(http.StatusOK, dtos.NewUserOutputDto(user))
}
