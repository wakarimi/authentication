package token_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Validate(c *gin.Context) {
	log.Debug().Msg("Validating a token")

	log.Debug().Msg("Tokens validated")
}
