package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	DatabaseConfiguration
}

type DatabaseConfiguration struct {
	DatabaseConnectionString string
}

func LoadConfiguration() (config *Configuration, err error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config = &Configuration{
		DatabaseConfiguration{
			DatabaseConnectionString: viper.GetString("WAKARIMI_AUTHENTICATION_DB_STRING"),
		},
	}

	return config, nil
}
