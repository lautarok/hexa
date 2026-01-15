package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/middlewares"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
	"github.com/lautarok/hexa/src/application/usecases"
)

type RolesController struct {
	getRolesUsecase   *usecases.GetRolesUsecase
	createRoleUsecase *usecases.CreateRoleUsecase
	authMiddleware    *middlewares.AuthMiddleware
	validation        ports.ValidationPort
}

type RolesControllerDeps struct {
	GetRolesUsecase   *usecases.GetRolesUsecase
	CreateRoleUsecase *usecases.CreateRoleUsecase
	AuthMiddleware    *middlewares.AuthMiddleware
	Validation        ports.ValidationPort
}

func NewRolesController(deps *RolesControllerDeps) *RolesController {
	return &RolesController{
		authMiddleware:    deps.AuthMiddleware,
		validation:        deps.Validation,
		getRolesUsecase:   deps.GetRolesUsecase,
		createRoleUsecase: deps.CreateRoleUsecase,
	}
}

func (controller *RolesController) Name() string {
	return "roles"
}

func (controller *RolesController) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.GET("/", controller.GetRoleList)
	group.POST("/", controller.CreateRole)
}

func (controller *RolesController) GetRoleList(ctx *gin.Context) {
	var paginationDto dtos.PaginationInputDto
	ctx.ShouldBindQuery(&paginationDto)
	if err := paginationDto.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	roles, appErr := controller.getRolesUsecase.GetRoleList(ctx, &usecases.GetRolesUsecaseInput{
		Page:  paginationDto.Page,
		Limit: paginationDto.Limit,
	})

	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	rolesDto := []*dtos.RoleOutputDto{}

	for _, role := range roles {
		var rolePermissions []*dtos.PermissionOutputDto
		for _, permission := range role.Permissions {
			rolePermissions = append(rolePermissions, &dtos.PermissionOutputDto{
				ID:        permission.ID,
				Alias:     permission.Alias,
				CreatedAt: permission.CreatedAt,
			})
		}

		rolesDto = append(rolesDto, &dtos.RoleOutputDto{
			ID:          role.ID,
			NameEn:      role.NameEn,
			NameEs:      role.NameEs,
			NameFr:      role.NameFr,
			NamePt:      role.NamePt,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
			Permissions: rolePermissions,
		})
	}

	ctx.JSON(http.StatusOK, &dtos.RoleListOutputDto{
		Roles: rolesDto,
		PaginationOutputDto: &dtos.PaginationOutputDto{
			Page:  paginationDto.Page,
			Limit: paginationDto.Limit,
		},
	})
}

func (controller *RolesController) CreateRole(ctx *gin.Context) {
	var reqBody dtos.CreateRoleInputDto
	ctx.ShouldBindBodyWithJSON(&reqBody)
	if err := reqBody.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	permissions := []*usecases.CreateRoleUsecaseInputPermission{}
	for _, permission := range reqBody.Permissions {
		permissions = append(permissions, &usecases.CreateRoleUsecaseInputPermission{
			Alias: permission.Alias,
		})
	}

	output, appErr := controller.createRoleUsecase.CreateRole(ctx, &usecases.CreateRoleUsecaseInput{
		NameEn:      reqBody.NameEn,
		NameEs:      reqBody.NameEs,
		NameFr:      reqBody.NameFr,
		NamePt:      reqBody.NamePt,
		Permissions: permissions,
	})
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	var rolePermissions []*dtos.PermissionOutputDto
	for _, permission := range output.Role.Permissions {
		rolePermissions = append(rolePermissions, &dtos.PermissionOutputDto{
			ID:        permission.ID,
			Alias:     permission.Alias,
			CreatedAt: permission.CreatedAt,
		})
	}

	ctx.JSON(http.StatusCreated, &dtos.RoleOutputDto{
		ID:          output.Role.ID,
		NameEn:      output.Role.NameEn,
		NameEs:      output.Role.NameEs,
		NameFr:      output.Role.NameFr,
		NamePt:      output.Role.NamePt,
		CreatedAt:   output.Role.CreatedAt,
		UpdatedAt:   output.Role.UpdatedAt,
		Permissions: rolePermissions,
	})
}
