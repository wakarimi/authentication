package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetVersion() string {
	return "v1"
}

func (h *Handler) AddRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-out", h.SignOut)
		auth.POST("/sign-out-all", h.SignOutAll)
	}
	tokens := api.Group("/tokens")
	{
		tokens.POST("/refresh", h.RefreshToken)
		tokens.POST("/verify", h.VerifyToken)
	}
	accounts := api.Group("/accounts")
	{
		accounts.GET("/me", h.GetRequestersAccount)
		accounts.GET("", h.GetAccounts)
		accounts.POST("/change-password", h.ChangePassword)
		accounts.POST("/sign-up", h.SignUp)

		account := accounts.Group("/:accountId")
		{
			roles := account.Group("/roles")
			{
				roles.POST("", h.AssignRole)
				roles.DELETE("", h.RevokeRole)
			}
		}
	}
}
