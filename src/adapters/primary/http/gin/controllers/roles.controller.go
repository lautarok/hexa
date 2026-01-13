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
	getRolesUsecase *usecases.GetRolesUsecase
	authMiddleware  *middlewares.AuthMiddleware
	validation      ports.ValidationPort
}

type RolesControllerDeps struct {
	GetRolesUsecase *usecases.GetRolesUsecase
	AuthMiddleware  *middlewares.AuthMiddleware
	Validation      ports.ValidationPort
}

func NewRolesController(deps *RolesControllerDeps) *RolesController {
	return &RolesController{
		authMiddleware:  deps.AuthMiddleware,
		validation:      deps.Validation,
		getRolesUsecase: deps.GetRolesUsecase,
	}
}

func (controller *RolesController) Name() string {
	return "roles"
}

func (controller *RolesController) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.GET("/", controller.GetRoleList)
}

func (controller *RolesController) GetRoleList(ctx *gin.Context) {
	var reqBody dtos.PaginationInputDto
	ctx.ShouldBindBodyWithJSON(&reqBody)
	if err := reqBody.Validate(controller.validation); err != nil {
		ctx.Error(
			errors.NewInvalidInputError(err.Error()),
		)
		return
	}

	roles, appErr := controller.getRolesUsecase.GetRoleList(ctx, &usecases.GetRolesInput{
		Page:  reqBody.Page,
		Limit: reqBody.Limit,
	})

	if appErr != nil {
		ctx.Error(appErr)
		return
	}

	rolesDto := []*dtos.RoleOutputDto{}

	for _, role := range roles {
		rolesDto = append(rolesDto, &dtos.RoleOutputDto{
			ID:        role.ID,
			NameEn:    role.NameEn,
			NameEs:    role.NameEs,
			NameFr:    role.NameFr,
			NamePt:    role.NamePt,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, &dtos.RoleListOutputDto{
		Roles: rolesDto,
		PaginationOutputDto: &dtos.PaginationOutputDto{
			Page:  reqBody.Page,
			Limit: reqBody.Limit,
		},
	})
}
