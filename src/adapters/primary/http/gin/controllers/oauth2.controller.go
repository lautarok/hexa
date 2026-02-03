package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/application/usecases/oauth2/command"
	"github.com/lautarok/hexa/src/application/usecases/oauth2/query"
	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/ports"
)

type OAuth2Controller struct {
	getGoogleSignOnUrlUsecase *query.GetGoogleSignOnURLUsecase
	googleSignOnUsecase       *command.GoogleSignOnUsecase
	validation                ports.ValidationPort
}

type OAuth2ControllerDeps struct {
	GetGoogleSignOnURLUsecase *query.GetGoogleSignOnURLUsecase
	GoogleSignOnUsecase       *command.GoogleSignOnUsecase
	Validation                ports.ValidationPort
}

func NewOAuth2Controller(deps *OAuth2ControllerDeps) *OAuth2Controller {
	return &OAuth2Controller{
		getGoogleSignOnUrlUsecase: deps.GetGoogleSignOnURLUsecase,
		googleSignOnUsecase:       deps.GoogleSignOnUsecase,
		validation:                deps.Validation,
	}
}

func (controller *OAuth2Controller) Name() string {
	return "oauth2"
}

func (controller *OAuth2Controller) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.GET("google-url", controller.GetGoogleURL)
	group.POST("google-sign-on", controller.GoogleSignOn)
}

func (controller *OAuth2Controller) GetGoogleURL(ctx *gin.Context) {
	type localeDto struct {
		Locale string `form:"locale" validate:"required"`
	}
	queryDto := localeDto{}
	ctx.ShouldBindQuery(&queryDto)
	if err := controller.validation.Struct(&queryDto); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}
	googleAuthUrl := controller.getGoogleSignOnUrlUsecase.GetGoogleAuthURL(ctx, queryDto.Locale)
	ctx.JSON(http.StatusOK, &dtos.URLOutputDto{
		URL: googleAuthUrl,
	})
}

func (controller *OAuth2Controller) GoogleSignOn(ctx *gin.Context) {
	type codeDto struct {
		Code string `json:"code" validate:"required"`
	}
	bodyDto := codeDto{}
	ctx.ShouldBindJSON(&bodyDto)
	if err := controller.validation.Struct(&bodyDto); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	type localeDto struct {
		Locale string `form:"locale" validate:"required"`
	}
	queryDto := localeDto{}
	ctx.ShouldBindQuery(&queryDto)
	if err := controller.validation.Struct(&queryDto); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	googleSignOn, appErr := controller.googleSignOnUsecase.GoogleSignOn(
		ctx, queryDto.Locale, bodyDto.Code,
	)
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(http.StatusOK, dtos.NewTokenOutputDto(
		googleSignOn.Token,
		googleSignOn.Exp,
		googleSignOn.User,
	))
}
