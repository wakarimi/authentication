package api

import (
	"authentication/internal/config"
	"authentication/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(cfg *config.Configuration) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/auth-service")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.POST("/register", handlers.Register)
		api.POST("/login", func(c *gin.Context) { handlers.Login(c, cfg) })
		api.POST("/refresh", func(c *gin.Context) { handlers.Refresh(c, cfg) })
		api.POST("/validate", func(c *gin.Context) { handlers.ValidateToken(c, cfg) })
	}

	return r
}
