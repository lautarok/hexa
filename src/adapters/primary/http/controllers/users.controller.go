package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/app/usecases"
)

type UsersController struct {
	getUsersUsecase *usecases.GetUsersUsecase
}

type UsersControllerDeps struct {
	GetUsersUsecase *usecases.GetUsersUsecase
}

func NewUsersController(deps *UsersControllerDeps) *UsersController {
	return &UsersController{
		getUsersUsecase: deps.GetUsersUsecase,
	}
}

func (controller *UsersController) Register(router any) {
	controllerName := "users"

	ginRouter, ok := router.(gin.IRouter)
	if !ok {
		log.Println("Controller [" + strings.ToUpper(controllerName) + "] only works with Gin")
		return
	}

	group := ginRouter.Group(controllerName)
	group.GET("/", controller.GetUserList)
}

func (controller *UsersController) GetUserList(ctx *gin.Context) {
	users, err := controller.getUsersUsecase.GetUserList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, users)
}
