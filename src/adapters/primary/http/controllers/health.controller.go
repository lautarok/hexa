package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (controller *HealthController) Register(router any) {
	controllerName := "health"

	ginRouter, ok := router.(gin.IRouter)
	if !ok {
		log.Println("Controller [" + strings.ToUpper(controllerName) + "] only works with Gin")
		return
	}

	group := ginRouter.Group(controllerName)
	group.GET("", controller.GetHealth)
}

func (controller *HealthController) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{
		"statusCode":     http.StatusOK,
		"health":         "100%",
		"responseTimeMs": 0,
	})
}
