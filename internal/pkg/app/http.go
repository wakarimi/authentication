package app

import (
	"fmt"
	"wakarimi-authentication/internal/config"
)

func (a *App) RegisterRoutes() {
	api := a.router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-in", a.handler.SignIn)
			auth.POST("/sign-out", a.handler.SignOut)
			auth.POST("/sign-out-all", a.handler.SignOutAll)
		}
		tokens := api.Group("/tokens")
		{
			tokens.POST("/refresh", a.handler.RefreshTokens)
			tokens.POST("/verify", a.handler.VerifyAccessToken)
		}
		accounts := api.Group("/accounts")
		{
			accounts.GET("/me", a.handler.GetRequestersAccount)
			accounts.GET("", a.handler.GetAccounts)
			accounts.POST("/change-password", a.handler.ChangePassword)
			accounts.POST("/sign-up", a.handler.SignUp)

			account := accounts.Group("/:accountId")
			{
				roles := account.Group("/roles")
				{
					roles.POST("", a.handler.AssignRole)
					roles.DELETE("", a.handler.RevokeRole)
				}
			}
		}
	}
}

func (a *App) StartHTTPServer(cfg config.HTTPConfig) error {
	err := a.router.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return err
	}
	return nil
}
