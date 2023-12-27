package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/config"
)

type Server struct {
	engine *gin.Engine
	host   string
	port   string
}

func NewServer(cfg config.HTTPConfig, engine *gin.Engine) *Server {
	gin.SetMode(gin.ReleaseMode)
	server := &Server{
		engine: engine,
		host:   cfg.Host,
		port:   cfg.Port,
	}
	return server
}

func (s *Server) RegisterRoutes(routers ...*Router) {
	for _, router := range routers {
		s.engine.Use(router.Engine().Handlers...)
	}
}

func (s *Server) Start() error {
	err := s.engine.Run(":" + s.port)
	if err != nil {
		log.Error().Err(err).Msg("Failed to start server")
		return err
	}

	return nil
}
