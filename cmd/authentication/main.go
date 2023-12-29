package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"wakarimi-authentication/internal/config"
	"wakarimi-authentication/internal/pkg/app"
)

func main() {
	configFilePath := flag.String("config", "config/config.yml", "Path to the config file")
	flag.Parse()

	cfg, err := config.New(*configFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config")
	}

	initializeLogger(cfg.App.LoggingLevel)

	application, err := app.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create application")
	}

	err = application.StartHTTPServer(cfg.HTTP)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}

func initializeLogger(level zerolog.Level) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().Caller().Logger().
		With().Str("service", "authentication").Logger().
		Level(level)
	log.Debug().Msg("Logger initialized")
}
