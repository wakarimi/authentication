package config

import (
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {
	Name             string
	Env              string
	Version          string
	RefreshSecretKey string
	AccessSecretKey  string
}

type HTTPConfig struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type ThirdPartyService struct {
	Url string
}

type DBConfig struct {
	Host         string
	Port         string
	DBName       string
	User         string
	Password     string
	Timeout      time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Charset      string
}

type Config struct {
	App        AppConfig
	HTTP       HTTPConfig
	DB         DBConfig
	ApiGateway ThirdPartyService
}

func NewConfig(filePath string) (config Config, err error) {
	viper.SetConfigFile(filePath)

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	viper.SetDefault("app.name", "wakarimi-authentication")
	viper.SetDefault("app.env", "dev")
	viper.SetDefault("app.version", "v1")
	viper.SetDefault("http.read_timeout", "1s")
	viper.SetDefault("http.write_timeout", "1s")
	viper.SetDefault("db.read_timeout", "1s")
	viper.SetDefault("db.write_timeout", "1s")
	viper.SetDefault("db.charset", "UTF-8")

	config = Config{
		App: AppConfig{
			Name:             viper.GetString("app.name"),
			Env:              viper.GetString("app.env"),
			Version:          viper.GetString("app.version"),
			RefreshSecretKey: viper.GetString("app.refresh_key"),
			AccessSecretKey:  viper.GetString("app.access_key"),
		},

		HTTP: HTTPConfig{
			Host:         viper.GetString("http.host"),
			Port:         viper.GetString("http.port"),
			ReadTimeout:  viper.GetDuration("http.read_timeout"),
			WriteTimeout: viper.GetDuration("http.write_timeout"),
		},

		DB: DBConfig{
			Host:         viper.GetString("db.host"),
			Port:         viper.GetString("db.port"),
			DBName:       viper.GetString("db.name"),
			User:         viper.GetString("db.user"),
			Password:     viper.GetString("db.password"),
			ReadTimeout:  viper.GetDuration("db.read_timeout"),
			WriteTimeout: viper.GetDuration("db.write_timeout"),
			Charset:      viper.GetString("db.charset"),
		},
	}

	return config, err
}
