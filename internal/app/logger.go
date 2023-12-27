package app

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"wakarimi-authentication/internal/config"
)

func (a *App) initLogger(cfg config.AppConfig) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().Caller().Logger().
		With().Str("service", "authentication").Logger().
		Level(cfg.LoggingLevel)
	log.Debug().Msg("Logger initialized")
	a.logger = log.Logger
}
