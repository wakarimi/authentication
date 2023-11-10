package api

import (
	"authentication/internal/context"
	"authentication/internal/database/repository/account_repo"
	"authentication/internal/database/repository/account_role_repo"
	"authentication/internal/database/repository/device_repo"
	"authentication/internal/handlers/account_handler"
	"authentication/internal/handlers/token_handler"
	"authentication/internal/middleware"
	"authentication/internal/service"
	"authentication/internal/service/account_role_service"
	"authentication/internal/service/account_service"
	"authentication/internal/service/device_service"
	"authentication/internal/service/token_service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/text/language"
)

func SetupRouter(ac *context.AppContext) (r *gin.Engine) {
	log.Debug().Msg("Router setup")
	gin.SetMode(gin.ReleaseMode)

	r = gin.New()
	r.Use(middleware.ZerologMiddleware(log.Logger))

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", json.Unmarshal)
	bundle.LoadMessageFile("internal/locales/en-US.json")
	bundle.LoadMessageFile("internal/locales/ru-RU.json")

	accountRepo := account_repo.NewRepository()
	accountRoleRepo := account_role_repo.NewRepository()
	deviceRepo := device_repo.NewRepository()

	txManager := service.NewTransactionManager(*ac.Db)

	accountRoleService := account_role_service.NewService(accountRoleRepo)
	accountService := account_service.NewService(accountRepo, *accountRoleService)
	deviceService := device_service.NewService(deviceRepo)
	tokenService := token_service.NewService(*accountService, *accountRoleService, *deviceService)

	accountHandler := account_handler.NewHandler(*accountService, *accountRoleService, txManager, bundle)
	tokenHandler := token_handler.NewHandler(*tokenService, txManager, bundle)

	api := r.Group("/api/auth-service")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		api.POST("/register", accountHandler.Create)

		api.POST("/login", tokenHandler.Create)

		token := api.Group("/token")
		{
			token.POST("/refresh", tokenHandler.Refresh)
			token.POST("/validate", tokenHandler.Validate)
		}
	}

	log.Debug().Msg("Router setup successfully")
	return r
}
