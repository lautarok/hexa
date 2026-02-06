package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/middlewares"
	permissionsQuery "github.com/lautarok/hexa/src/application/usecases/permissions/query"
	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/ports"
)

type PermissionsController struct {
	getPermissionsUsecase *permissionsQuery.GetPermissionsUsecase
	authMiddleware        *middlewares.AuthMiddleware
	validation            ports.ValidationPort
}

type PermissionsControllerDeps struct {
	GetPermissionsUsecase *permissionsQuery.GetPermissionsUsecase
	AuthMiddleware        *middlewares.AuthMiddleware
	Validation            ports.ValidationPort
}

func NewPermissionsController(deps *PermissionsControllerDeps) *PermissionsController {
	return &PermissionsController{
		getPermissionsUsecase: deps.GetPermissionsUsecase,
		authMiddleware:        deps.AuthMiddleware,
		validation:            deps.Validation,
	}
}

func (controller *PermissionsController) Name() string {
	return "permissions"
}

func (controller *PermissionsController) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.GET("/", controller.authMiddleware.HandleAuth("admin"), controller.GetPermissionList)
}

func (controller *PermissionsController) GetPermissionList(ctx *gin.Context) {
	var paginationDto dtos.PaginationInputDto
	ctx.ShouldBindQuery(&paginationDto)
	if err := paginationDto.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	result, appErr := controller.getPermissionsUsecase.GetPermissions(ctx, &permissionsQuery.GetPermissionsUsecaseInput{
		Page:  paginationDto.Page,
		Limit: paginationDto.Limit,
	})
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(http.StatusOK, dtos.NewPermissionListOutputDto(
		&paginationDto,
		result,
	))
}
