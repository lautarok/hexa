package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/app/shared/ports"
)

type GinAdapter struct {
	Engine *gin.Engine
	Router gin.IRouter
}

func NewGinAdapter(routerPrefix string, prodEnv bool) ports.HTTPPort {
	if prodEnv {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}

	engine := gin.Default()

	return &GinAdapter{
		Engine: engine,
		Router: engine.Group(routerPrefix),
	}
}

func (adapter *GinAdapter) Start(addr string) error {
	server := adapter.Engine
	log.Println("Gin server is listening on address: \"" + addr + "\"")
	err := server.Run(addr)
	return err
}

func (adapter *GinAdapter) RegisterControllers(controllers ...ports.HTTPRouteRegister) {
	for _, controller := range controllers {
		controller.Register(adapter.Router)
	}
}
