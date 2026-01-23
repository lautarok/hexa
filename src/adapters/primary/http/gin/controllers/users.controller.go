package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
	usersQuery "github.com/lautarok/hexa/src/application/usecases/users/query"
)

type UsersController struct {
	getUsersUsecase *usersQuery.GetUsersUsecase
	validation      ports.ValidationPort
}

type UsersControllerDeps struct {
	GetUsersUsecase *usersQuery.GetUsersUsecase
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
	var paginationDto dtos.PaginationInputDto
	ctx.ShouldBindQuery(&paginationDto)
	if err := paginationDto.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	users, appErr := controller.getUsersUsecase.GetUserList(ctx, &usersQuery.GetUsersUsecaseInput{
		Page:  paginationDto.Page,
		Limit: paginationDto.Limit,
	})
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(http.StatusOK, dtos.NewUserListOutputDto(&paginationDto, users))
}
