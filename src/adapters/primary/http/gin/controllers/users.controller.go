package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
	"github.com/lautarok/hexa/src/application/usecases"
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
	var paginationDto dtos.PaginationInputDto
	ctx.ShouldBindQuery(&paginationDto)
	if err := paginationDto.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	users, appErr := controller.getUsersUsecase.GetUserList(ctx, &usecases.GetUsersInput{
		Page:  paginationDto.Page,
		Limit: paginationDto.Limit,
	})
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	usersDto := []*dtos.UserOutputDto{}

	for _, user := range users {
		usersDto = append(usersDto, &dtos.UserOutputDto{
			ID:      user.ID,
			Name:    user.Name,
			Surname: user.Surname,
			Credential: &dtos.CredentialOutputDto{
				ID:       user.Credential.ID,
				Email:    user.Credential.Email,
				Username: user.Credential.Username,
			},
			CreatedAt: user.CreatedAt,
		})
	}

	output := dtos.UserListOutputDto{
		PaginationOutputDto: &dtos.PaginationOutputDto{
			Page:  paginationDto.Page,
			Limit: paginationDto.Limit,
		},
		Users: usersDto,
	}

	ctx.JSON(http.StatusOK, output)
}
