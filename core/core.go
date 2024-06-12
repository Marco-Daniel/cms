package core

import (
	"cms/config"
	"cms/routes"

	"github.com/gin-gonic/gin"
)

type Config struct {
	DBURI string
}

type Server struct {
	Engine *gin.Engine
}

func CreateServer(cfg Config) (*Server, error) {
	err := config.ConnectDB(cfg.DBURI)
	if err != nil {
		return nil, err
	}

	engine := gin.Default()

	return &Server{Engine: engine}, nil
}

func (server *Server) RegiserRoutes(routeFuncs []func(*gin.Engine)) {
	for _, routeFunc := range routeFuncs {
		routeFunc(server.Engine)
	}
	routes.RegisterSignupRoute(server.Engine)
}

func (server *Server) AddMiddleware(middlewareFuncs []gin.HandlerFunc) {
	for _, middlewareFunc := range middlewareFuncs {
		server.Engine.Use(middlewareFunc)
	}
}

func (server *Server) Start() {
	server.Engine.Run("localhost:8080")
}
