package app

import (
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/handler/http"
	v1 "wakarimi-authentication/internal/handler/http/api/v1"
)

func (a *App) StartHTTPServer() {
	handler := v1.NewHandler(a.container.GetUseCase(), a.logger)

	a.router.
		WithSwagger().
		WithHandler(handler, a.logger)

	server := http.NewServer(a.cfg.HTTP, a.router.Engine())
	server.RegisterRoutes(a.router)

	err := server.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start http server")
	}
}
