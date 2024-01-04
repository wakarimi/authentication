package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"wakarimi-authentication/internal/handler/response"
)

func (h *Handler) SignOutAll(c *gin.Context) {
	log.Debug().Msg("Sign out for account")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(&h.bundle, lang)

	accountIDHeader := c.GetHeader("X-Account-ID")
	requesterAccountID, err := strconv.Atoi(accountIDHeader)
	if err != nil {
		log.Error().Err(err).Str("accountIdHeader", accountIDHeader).Msg("Invalid X-Account-ID format")
		c.JSON(http.StatusForbidden, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "InvalidHeaderFormat",
			}),
			Reason: err.Error(),
		})
		return
	}

	err = h.useCase.SignOutAll(requesterAccountID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to sign out by account")
		messageID := "FailedToSignOut"
		message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
		if errLoc != nil {
			message = h.engLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
		}
		c.JSON(http.StatusInternalServerError, response.Error{
			Message: message,
			Reason:  err.Error(),
		})
		return
	}

	log.Debug().Msg("Refresh token for account deleted")
	c.Status(http.StatusNoContent)
}
