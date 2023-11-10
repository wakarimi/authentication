package account_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Create(c *gin.Context) {
	log.Debug().Msg("Creating an account")

	log.Debug().Msg("Account created")
}
