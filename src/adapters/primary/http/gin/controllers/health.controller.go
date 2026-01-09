package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (controller *HealthController) Name() string {
	return "health"
}

func (controller *HealthController) Register(router *gin.RouterGroup) {
	group := router.Group(controller.Name())
	group.GET("", controller.GetHealth)
}

func (controller *HealthController) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{
		"statusCode":     http.StatusOK,
		"health":         "100%",
		"responseTimeMs": 0,
	})
}
