package health

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthService  *HealthService
	controllerName string
}

type HealthControllerDeps struct {
	HealthService *HealthService
}

func NewHealthController(deps *HealthControllerDeps) *HealthController {
	return &HealthController{
		healthService: deps.HealthService,
	}
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
	health, err := controller.healthService.GetHealth()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Internal server error",
		})
	}

	ctx.JSON(http.StatusOK, health)
}
