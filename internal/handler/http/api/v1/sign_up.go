package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	a, err := h.uc.SignUp(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(201, a)
}
