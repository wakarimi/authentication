package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type HandlerRouter interface {
	GetVersion() string
	AddRoutes(r *gin.RouterGroup)
}

type Router struct {
	engine *gin.Engine
}

func NewRouter() *Router {
	gin.SetMode(gin.ReleaseMode)
	return &Router{
		engine: gin.New(),
	}
}

func (r *Router) WithSwagger() *Router {
	log.Warn().Msg("Implement swagger endpoint")
	return r
}

func (r *Router) WithHandler(h HandlerRouter, logger zerolog.Logger) *Router {
	api := r.engine.Group("/api/" + h.GetVersion())

	h.AddRoutes(api)

	return r
}

func (r *Router) Engine() *gin.Engine {
	return r.engine
}
