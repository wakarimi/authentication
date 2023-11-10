package token_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Refresh(c *gin.Context) {
	log.Debug().Msg("Refreshing tokens")

	log.Debug().Msg("Tokens refreshed")
}
