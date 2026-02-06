package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/middlewares"
	"github.com/lautarok/hexa/src/application/usecases/users/command"
	"github.com/lautarok/hexa/src/application/usecases/users/query"
	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/ports"
)

type UsersController struct {
	getUsersUsecase         *query.GetUsersUsecase
	updateCredentialUsecase *command.UpdateCredentialUsecase
	authMiddleware          *middlewares.AuthMiddleware
	validation              ports.ValidationPort
}

type UsersControllerDeps struct {
	GetUsersUsecase         *query.GetUsersUsecase
	UpdateCredentialUsecase *command.UpdateCredentialUsecase
	AuthMiddleware          *middlewares.AuthMiddleware
	Validation              ports.ValidationPort
}

func NewUsersController(deps *UsersControllerDeps) *UsersController {
	return &UsersController{
		getUsersUsecase:         deps.GetUsersUsecase,
		updateCredentialUsecase: deps.UpdateCredentialUsecase,
		authMiddleware:          deps.AuthMiddleware,
		validation:              deps.Validation,
	}
}

func (controller *UsersController) Name() string {
	return "users"
}

func (controller *UsersController) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.GET("/", controller.authMiddleware.HandleAuth("admin"), controller.GetUserList)
	group.PUT("/credential", controller.authMiddleware.HandleAuth("user:self", "admin"), controller.UpdateCredential)
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

	users, appErr := controller.getUsersUsecase.GetUserList(ctx, &query.GetUsersUsecaseInput{
		Page:  paginationDto.Page,
		Limit: paginationDto.Limit,
	})
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(http.StatusOK, dtos.NewUserListOutputDto(&paginationDto, users))
}

func (controller *UsersController) UpdateCredential(ctx *gin.Context) {
	updateDto := dtos.UpdateUserCredentialInputDto{}
	err := ctx.ShouldBindJSON(&updateDto)
	if err != nil {
		ctx.Error(err)
		return
	} else if err := updateDto.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	credential, appErr := controller.updateCredentialUsecase.UpdateCredential(ctx, &command.UpdateCredentialUsecaseInput{
		UserID:   updateDto.UserID,
		Username: updateDto.Username,
	})
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(http.StatusOK, &dtos.CredentialOutputDto{
		ID:       credential.ID,
		Username: credential.Username,
		Email:    credential.Email,
	})
}
