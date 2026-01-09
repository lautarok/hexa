package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/app/dtos"
	"github.com/lautarok/hexa/src/app/errors"
	"github.com/lautarok/hexa/src/app/ports"
	"github.com/lautarok/hexa/src/app/usecases"
)

type UsersController struct {
	getUsersUsecase *usecases.GetUsersUsecase
	validation      ports.ValidationPort
}

type UsersControllerDeps struct {
	GetUsersUsecase *usecases.GetUsersUsecase
	Validation      ports.ValidationPort
}

func NewUsersController(deps *UsersControllerDeps) *UsersController {
	return &UsersController{
		getUsersUsecase: deps.GetUsersUsecase,
		validation:      deps.Validation,
	}
}

func (controller *UsersController) Name() string {
	return "users"
}

func (controller *UsersController) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.GET("/", controller.GetUserList)
}

func (controller *UsersController) GetUserList(ctx *gin.Context) {
	var paginationDto dtos.PaginationInput
	err := ctx.ShouldBindQuery(&paginationDto)
	if err != nil || (paginationDto.Page == 0 && paginationDto.Limit == 0) {
		paginationDto = dtos.PaginationInput{
			Page:  1,
			Limit: 15,
		}
	} else if err = controller.validation.Struct(paginationDto); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(strings.Split(err.Error(), "\n")[0]),
		)
		return
	}

	users, err := controller.getUsersUsecase.GetUserList(ctx, &paginationDto)
	if err != nil {
		ctx.Error(
			errors.NewInternalError(err),
		)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
