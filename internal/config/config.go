package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	DatabaseConfiguration
	HttpServerConfiguration
}

type DatabaseConfiguration struct {
	DatabaseConnectionString string
}

type HttpServerConfiguration struct {
	Port             string
	MicroservicesIps []string
}

func LoadConfiguration() (config *Configuration, err error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	localIps := []string{"127.0.0.1", "::1"}
	config = &Configuration{
		DatabaseConfiguration{
			DatabaseConnectionString: viper.GetString("WAKARIMI_AUTHENTICATION_DB_STRING"),
		},
		HttpServerConfiguration{
			Port:             viper.GetString("HTTP_SERVER_PORT"),
			MicroservicesIps: append(localIps, parseIp4s(viper.GetString("ALLOWED_IP4"))...),
		},
	}

	return config, nil
}

func parseIp4s(ips string) []string {
	parts := strings.Split(ips, "_")
	var result []string
	for i := 0; i < len(parts); i += 4 {
		if i+3 < len(parts) {
			result = append(result, fmt.Sprintf("%s.%s.%s.%s", parts[i], parts[i+1], parts[i+2], parts[i+3]))
		}
	}
	return result
}
