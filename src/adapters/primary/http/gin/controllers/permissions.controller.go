package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/core/errors"
	"github.com/lautarok/hexa/src/core/ports"
	permissionsQuery "github.com/lautarok/hexa/src/core/usecases/permissions/query"
)

type PermissionsController struct {
	getPermissionsUsecase *permissionsQuery.GetPermissionsUsecase
	validation            ports.ValidationPort
}

type PermissionsControllerDeps struct {
	GetPermissionsUsecase *permissionsQuery.GetPermissionsUsecase
	Validation            ports.ValidationPort
}

func NewPermissionsController(deps *PermissionsControllerDeps) *PermissionsController {
	return &PermissionsController{
		getPermissionsUsecase: deps.GetPermissionsUsecase,
		validation:            deps.Validation,
	}
}

func (controller *PermissionsController) Name() string {
	return "permissions"
}

func (controller *PermissionsController) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.GET("/", controller.GetPermissionList)
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
