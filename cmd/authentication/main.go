package main

import (
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/app"
)

var configFilePath string

func main() {
	configFilePath = "./config/config.yml"

	a, err := app.NewApp(configFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create app")
	}

	a.StartHTTPServer()
}
