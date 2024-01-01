package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
	"wakarimi-authentication/internal/config"
	"wakarimi-authentication/internal/handler"
	"wakarimi-authentication/internal/middleware"
	"wakarimi-authentication/internal/service"
	"wakarimi-authentication/internal/service/access_token_service"
	"wakarimi-authentication/internal/service/account_role_service"
	"wakarimi-authentication/internal/service/account_service"
	"wakarimi-authentication/internal/service/device_service"
	"wakarimi-authentication/internal/service/refresh_token_service"
	"wakarimi-authentication/internal/storage"
	"wakarimi-authentication/internal/storage/repo/account_repo"
	"wakarimi-authentication/internal/storage/repo/account_role_repo"
	"wakarimi-authentication/internal/storage/repo/device_repo"
	"wakarimi-authentication/internal/storage/repo/refresh_token_repo"
	"wakarimi-authentication/internal/use_case"
)

type App struct {
	handler *handler.Handler
	router  *gin.Engine
}

func New(cfg config.Config) (*App, error) {
	application := &App{}

	db, err := storage.New(cfg.DB)
	if err != nil {
		return nil, err
	}

	transactor := service.NewTransactor(*db)

	accountRepo := account_repo.New()
	accountRoleRepo := account_role_repo.New()
	deviceRepo := device_repo.New()
	refreshTokenRepo := refresh_token_repo.New()

	accessTokenService := access_token_service.New(cfg.App.AccessSecretKey)
	accountService := account_service.New(accountRepo)
	accountRoleService := account_role_service.New(accountRoleRepo)
	deviceService := device_service.New(deviceRepo)
	refreshTokenService := refresh_token_service.New(cfg.App.RefreshSecretKey, refreshTokenRepo)

	useCase := use_case.New(transactor, accessTokenService, accountService,
		accountRoleService, deviceService, refreshTokenService)

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", json.Unmarshal)
	bundle.LoadMessageFile("internal/locale/en-US.json")
	bundle.LoadMessageFile("internal/locale/ru-RU.json")
	application.handler = handler.New(useCase, *bundle)

	gin.SetMode(gin.ReleaseMode)
	application.router = gin.New()
	application.router.Use(middleware.ZerologMiddleware(log.Logger))
	application.router.Use(middleware.ProduceLanguageMiddleware())
	application.RegisterRoutes()

	return application, nil
}
