package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
	"github.com/lautarok/hexa/src/application/usecases"
	"github.com/lautarok/hexa/src/dtos"
)

type AuthController struct {
	signupUsecase *usecases.SignupUsecase
	validation    ports.ValidationPort
}

type AuthControllerDeps struct {
	SignupUsecase *usecases.SignupUsecase
	Validation    ports.ValidationPort
}

func NewAuthController(deps *AuthControllerDeps) *AuthController {
	return &AuthController{
		signupUsecase: deps.SignupUsecase,
		validation:    deps.Validation,
	}
}

func (controller *AuthController) Name() string {
	return "auth"
}

func (controller *AuthController) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.POST("signup", controller.Signup)
}

func (controller *AuthController) Signup(ctx *gin.Context) {
	var reqBody dtos.SignupInputDto
	if err := ctx.ShouldBindBodyWithJSON(&reqBody); err != nil {
		ctx.Error(
			errors.NewInvalidInputError("Wrong body"),
		)
		return
	}

	if err := controller.validation.Struct(reqBody); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(
				strings.Split(err.Error(), "\n")[0],
			),
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

	token, err := controller.signupUsecase.Signup(ctx, usecaseInput)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, &dtos.TokenOutputDto{
		Token: token,
	})
}
