package gin

import (
	"log"

	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	Engine *gin.Engine
	Router *gin.RouterGroup
}

func NewGinAdapter(routerPrefix string, prodEnv bool) *GinAdapter {
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

type Controller interface {
	Name() string
	Register(group *gin.RouterGroup)
}

func (adapter *GinAdapter) RegisterControllers(controllers ...Controller) {
	for _, controller := range controllers {
		controller.Register(adapter.Router)
	}
}

func (adapter *GinAdapter) RegisterGlobalMiddlewares(middlewares ...gin.HandlerFunc) {
	for _, middleware := range middlewares {
		adapter.Router.Use(middleware)
	}
}
