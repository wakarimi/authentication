package token_handler

import (
	"authentication/internal/errors"
	"authentication/internal/handler/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
)

// SignOut from the account
// @Summary SignOut from the account
// @Tags Tokens
// @Accept json
// @Produce json
// @Param Produce-Language header string false "Language preference" default(en-US)
// @Param X-Device-ID header int true "Device ID"
// @Success 204
// @Failure 400 {object} response.Error "Invalid X-Device-ID format"
// @Failure 404 {object} response.Error "Device not found; Token not found"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /auth/sign-out [post]
func (h *Handler) SignOut(c *gin.Context) {
	log.Debug().Msg("Sign out from the one device")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	deviceIDHeader := c.GetHeader("X-Device-ID")
	deviceID, err := strconv.Atoi(deviceIDHeader)
	if err != nil {
		log.Error().Err(err).Str("deviceIdHeader", deviceIDHeader).Msg("Invalid X-Device-ID format")
		c.JSON(http.StatusForbidden, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "InvalidHeaderFormat",
			}),
			Reason: err.Error(),
		})
		return
	}
	log.Debug().Int("deviceId", deviceID).Msg("Header X-Device-ID read successfully")

	err = h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = h.TokenService.DeleteByDevice(tx, deviceID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to sign out")
		if _, ok := err.(errors.NotFound); ok {
			messageID := "NotFound"
			message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
			if errLoc != nil {
				message = h.EngLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
			}
			c.JSON(http.StatusUnauthorized, response.Error{
				Message: message,
				Reason:  err.Error(),
			})
			return
		} else {
			messageID := "FailedToSignOut"
			message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
			if errLoc != nil {
				message = h.EngLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
			}
			c.JSON(http.StatusInternalServerError, response.Error{
				Message: message,
				Reason:  err.Error(),
			})
			return
		}
	}

	log.Debug().Msg("Sign out from account successfully")
	c.Status(http.StatusNoContent)
}
