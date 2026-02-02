package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/middlewares"
	"github.com/lautarok/hexa/src/core/errors"
	"github.com/lautarok/hexa/src/core/ports"
	rolesCommand "github.com/lautarok/hexa/src/core/usecases/roles/command"
	rolesQuery "github.com/lautarok/hexa/src/core/usecases/roles/query"
)

type RolesController struct {
	getRolesUsecase   *rolesQuery.GetRolesUsecase
	createRoleUsecase *rolesCommand.CreateRoleUsecase
	authMiddleware    *middlewares.AuthMiddleware
	validation        ports.ValidationPort
}

type RolesControllerDeps struct {
	GetRolesUsecase   *rolesQuery.GetRolesUsecase
	CreateRoleUsecase *rolesCommand.CreateRoleUsecase
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

	roles, appErr := controller.getRolesUsecase.GetRoleList(ctx, &rolesQuery.GetRolesUsecaseInput{
		Page:  paginationDto.Page,
		Limit: paginationDto.Limit,
	})

	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(http.StatusOK, dtos.NewRoleListOutputDto(
		&paginationDto,
		roles,
	))
}

func (controller *RolesController) CreateRole(ctx *gin.Context) {
	var reqBody dtos.CreateRoleInputDto
	ctx.ShouldBindJSON(&reqBody)
	if err := reqBody.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	output, appErr := controller.createRoleUsecase.CreateRole(ctx, reqBody.ToUsecaseInput())
	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	ctx.JSON(http.StatusCreated, dtos.NewRoleOutputDto(output))
}
