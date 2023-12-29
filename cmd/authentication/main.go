package main

import (
	"flag"
	"github.com/rs/zerolog/log"
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

	application, err := app.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create application")
	}

	err = application.StartHTTPServer(cfg.HTTP)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
