package config

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Configuration struct {
	Database
	HttpServer
	JwtConfiguration
	Logger
}

type Database struct {
	ConnectionString string
}

type HttpServer struct {
	Port string
}

type JwtConfiguration struct {
	RefreshSecretKey string
	AccessSecretKey  string
}

type Logger struct {
	Level zerolog.Level
}

func LoadConfiguration() (config *Configuration, err error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	dbConnectionString := viper.GetString("WAKARIMI_AUTHENTICATION_DB_STRING")
	if dbConnectionString == "" {
		err = fmt.Errorf("database connection string not found")
		log.Error().Err(err).Msg("Database connection string not found")
		return nil, err
	}

	httpPort := viper.GetString("HTTP_SERVER_PORT")
	if httpPort == "" {
		httpPort = "8020"
	}

	refreshSecretKey := viper.GetString("REFRESH_SECRET_KEY")
	if refreshSecretKey == "" {
		return nil, fmt.Errorf("refresh secret key not found")
	}

	accessSecretKey := viper.GetString("ACCESS_SECRET_KEY")
	if accessSecretKey == "" {
		return nil, fmt.Errorf("access secret key not found")
	}

	config = &Configuration{
		Database{
			ConnectionString: dbConnectionString,
		},
		HttpServer{
			Port: httpPort,
		},
		JwtConfiguration{
			RefreshSecretKey: refreshSecretKey,
			AccessSecretKey:  accessSecretKey,
		},
		Logger{
			Level: loadLoggingLevel(),
		},
	}

	return config, nil
}

func loadLoggingLevel() zerolog.Level {
	levelStr := viper.GetString("LOGGING_LEVEL")
	switch levelStr {
	case "TRACE":
		return zerolog.TraceLevel
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	case "WARN":
		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "FATAL":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
