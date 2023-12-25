package v1

import "github.com/gin-gonic/gin"

func (h *Handler) GetVersion() string {
	return "v1"
}

func (h *Handler) AddRoutes(r *gin.RouterGroup) {
	r.Group("/" + h.GetVersion())
}
