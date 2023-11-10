package token_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Create(c *gin.Context) {
	log.Debug().Msg("Creating a first refresh token")

	log.Debug().Msg("Refresh token generated")
}
