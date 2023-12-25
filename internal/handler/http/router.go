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
	router *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		router: gin.New(),
	}
}

func (r *Router) WithSwagger() *Router {
	log.Warn().Msg("Implement swagger endpoint")
	return r
}

func (r *Router) WithHandler(h HandlerRouter, logger zerolog.Logger) *Router {
	api := r.router.Group("/api")

	h.AddRoutes(api)

	return r
}
