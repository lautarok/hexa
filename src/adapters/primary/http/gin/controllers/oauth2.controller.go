package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/core/errors"
	"github.com/lautarok/hexa/src/core/ports"
	"github.com/lautarok/hexa/src/core/usecases/oauth2/query"
)

type OAuth2Controller struct {
	getGoogleOauth2UrlUsecase  *query.GetGoogleOAuth2URLUsecase
	getGoogleOauth2UserUsecase *query.GetGoogleOAuth2UserUsecase
	validation                 ports.ValidationPort
}

type OAuth2ControllerDeps struct {
	GetGoogleOAuth2UrlUsecase  *query.GetGoogleOAuth2URLUsecase
	GetGoogleOAuth2UserUsecase *query.GetGoogleOAuth2UserUsecase
	Validation                 ports.ValidationPort
}

func NewOAuth2Controller(deps *OAuth2ControllerDeps) *OAuth2Controller {
	return &OAuth2Controller{
		getGoogleOauth2UrlUsecase:  deps.GetGoogleOAuth2UrlUsecase,
		getGoogleOauth2UserUsecase: deps.GetGoogleOAuth2UserUsecase,
		validation:                 deps.Validation,
	}
}

func (controller *OAuth2Controller) Name() string {
	return "oauth2"
}

func (controller *OAuth2Controller) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.GET("google-url", controller.GetGoogleURL)
	group.GET("test", controller.GetGoogleCallback)
}

func (controller *OAuth2Controller) GetGoogleURL(ctx *gin.Context) {
	googleAuthUrl := controller.getGoogleOauth2UrlUsecase.GetGoogleAuthURL(ctx)
	ctx.String(http.StatusOK, googleAuthUrl)
}

func (controller *OAuth2Controller) GetGoogleCallback(ctx *gin.Context) {
	type codeDto struct {
		Code string `form:"code" validate:"required"`
	}
	queryDto := codeDto{}
	ctx.ShouldBindQuery(&queryDto)
	if err := controller.validation.Struct(&queryDto); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	googleCallback, appErr := controller.getGoogleOauth2UserUsecase.GetGoogleUser(ctx, queryDto.Code)
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(http.StatusOK, googleCallback)
}
