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

// SignOutAll from the account
// @Summary SignOutAll from the account
// @Tags Tokens
// @Accept json
// @Produce json
// @Param Produce-Language header string false "Language preference" default(en-US)
// @Param X-Account-ID header int true "Account ID"
// @Success 204
// @Failure 400 {object} response.Error "Invalid X-Account-ID format"
// @Failure 404 {object} response.Error "Account not found"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /auth/sign-out-all [post]
func (h *Handler) SignOutAll(c *gin.Context) {
	log.Debug().Msg("Sign out from the all devices")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	accountIDHeader := c.GetHeader("X-Account-ID")
	accountID, err := strconv.Atoi(accountIDHeader)
	if err != nil {
		log.Error().Err(err).Str("deviceIdHeader", accountIDHeader).Msg("Invalid X-Account-ID format")
		c.JSON(http.StatusForbidden, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "InvalidHeaderFormat",
			}),
			Reason: err.Error(),
		})
		return
	}
	log.Debug().Int("accountId", accountID).Msg("Header X-Account-ID read successfully")

	err = h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = h.TokenService.DeleteByAccount(tx, accountID)
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
