package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
	"github.com/lautarok/hexa/src/application/usecases"
)

type PermissionsController struct {
	getPermissionsUsecase *usecases.GetPermissionsUsecase
	validation            ports.ValidationPort
}

type PermissionsControllerDeps struct {
	GetPermissionsUsecase *usecases.GetPermissionsUsecase
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

	permissionList, appErr := controller.getPermissionsUsecase.GetPermissions(ctx, &usecases.GetPermissionsUsecaseInput{
		Page:  paginationDto.Page,
		Limit: paginationDto.Limit,
	})
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	permissionsDto := []*dtos.PermissionOutputDto{}
	for _, permission := range permissionList {
		permissionsDto = append(permissionsDto, &dtos.PermissionOutputDto{
			ID:        permission.ID,
			Alias:     permission.Alias,
			CreatedAt: permission.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, &dtos.PermissionListOutputDto{
		PaginationOutputDto: &dtos.PaginationOutputDto{
			Page:  paginationDto.Page,
			Limit: paginationDto.Limit,
		},
		Permissions: permissionsDto,
	})
}
