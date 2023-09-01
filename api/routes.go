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

	api := r.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		registerGroup := api.Group("/register") // Создаем группу для /register
		{
			registerGroup.POST("/user", handlers.RegisterUser)
			registerGroup.POST("/microservice", func(c *gin.Context) { handlers.RegisterMicroservice(c, cfg) })
		}
	}

	return r
}
