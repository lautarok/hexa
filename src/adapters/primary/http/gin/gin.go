package gin

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	Engine *gin.Engine
	Router *gin.RouterGroup
}

type GinAdapterDeps struct {
	RouterPrefix     string
	Production       bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           time.Duration
}

func NewGinAdapter(deps *GinAdapterDeps) *GinAdapter {
	if deps.Production {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}

	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     deps.AllowOrigins,
		AllowMethods:     deps.AllowMethods,
		AllowCredentials: deps.AllowCredentials,
		MaxAge:           deps.MaxAge,
	}))

	return &GinAdapter{
		Engine: engine,
		Router: engine.Group(deps.RouterPrefix),
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
