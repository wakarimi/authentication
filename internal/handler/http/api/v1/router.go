package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetVersion() string {
	return "v1"
}

func (h *Handler) AddRoutes(r *gin.RouterGroup) {
	r.GET("/a", func(context *gin.Context) {
		fmt.Print("Жопа")
	})
	apiV1 := r.Group("/" + h.GetVersion())
	{
		apiV1.POST("/sign-up", h.SignUp)
	}
}
