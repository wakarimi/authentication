package api

import (
	"authentication/internal/context"
	"authentication/internal/middleware"
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
	r.Use(middleware.ProduceLanguageMiddleware())

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", json.Unmarshal)
	bundle.LoadMessageFile("internal/locale/en-US.json")
	bundle.LoadMessageFile("internal/locale/ru-RU.json")

	// accountRepo := account_repo.NewRepository()
	// accountRoleRepo := account_role_repo.NewRepository()
	// deviceRepo := device_repo.NewRepository()
	// refreshTokenRepo := refresh_token_repo.NewRepository()

	// txManager := service.NewTransactionManager(*ac.Db)

	// accountRoleService := account_role_service.NewService(accountRoleRepo)
	// accountService := account_service.NewService(accountRepo, *accountRoleService)
	// deviceService := device_service.NewService(deviceRepo, *accountService)
	// tokenService := token_service.NewService(refreshTokenRepo, ac.Config.RefreshSecretKey, ac.Config.AccessSecretKey, *accountService, *accountRoleService, *deviceService)

	// accountHandler := account_handler.NewHandler(*accountService, *accountRoleService, txManager, bundle)
	// tokenHandler := token_handler.NewHandler(*tokenService, txManager, bundle)

	api := r.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		auth := api.Group("auth")
		{
			auth.POST("/sign-in")
			auth.POST("/sign-out")
			auth.POST("/sign-out-all")
		}

		tokens := api.Group("/tokens")
		{
			tokens.POST("/refresh")
			tokens.POST("/verify")
		}

		accounts := api.Group("accounts")
		{
			accounts.GET("/me")
			accounts.GET("")
			accounts.POST("/change-password")
			accounts.POST("/sign-up")

			account := accounts.Group(":accountId")
			{
				roles := account.Group("roles")
				{
					roles.POST("")
					roles.DELETE("")
				}
			}
		}
	}

	log.Debug().Msg("Router setup successfully")
	return r
}
