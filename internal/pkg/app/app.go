package app

import (
	"github.com/gin-gonic/gin"
	"wakarimi-authentication/internal/config"
	"wakarimi-authentication/internal/handler"
	"wakarimi-authentication/internal/service"
	"wakarimi-authentication/internal/service/access_token_service"
	"wakarimi-authentication/internal/service/account_role_service"
	"wakarimi-authentication/internal/service/account_service"
	"wakarimi-authentication/internal/service/device_service"
	"wakarimi-authentication/internal/service/refresh_token_service"
	"wakarimi-authentication/internal/storage"
	"wakarimi-authentication/internal/storage/account_repo"
	"wakarimi-authentication/internal/storage/account_role_repo"
	"wakarimi-authentication/internal/storage/device_repo"
	"wakarimi-authentication/internal/storage/refresh_token_repo"
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

	accessTokenService := access_token_service.New()
	accountService := account_service.New(accountRepo)
	accountRoleService := account_role_service.New(accountRoleRepo)
	deviceService := device_service.New(deviceRepo)
	refreshTokenService := refresh_token_service.New(refreshTokenRepo)

	useCase := use_case.New(transactor, accessTokenService, accountService,
		accountRoleService, deviceService, refreshTokenService)

	application.handler = handler.New(useCase)

	gin.SetMode(gin.ReleaseMode)
	application.router = gin.New()
	application.RegisterRoutes()

	return application, nil
}
