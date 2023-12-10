package account_handler

import (
	"authentication/internal/errors"
	"authentication/internal/handler/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
)

// changePasswordRequest represents the request body for password changing
type changePasswordRequest struct {
	// OldPassword for logging into your account
	OldPassword string `json:"oldPassword" validate:"required,alphanum"`
	// NewPassword for logging into your account
	NewPassword string `json:"newPassword" validate:"required,alphanum"`
}

// ChangePassword
// @Summary Change password
// @Tags Accounts
// @Accept json
// @Produce json
// @Param Produce-Language header string false "Language preference" default(en-US)
// @Param request body changePasswordRequest true "Old and new passwords"
// @Param X-Account-ID header int true "Account ID"
// @Success 204
// @Failure 400 {object} response.Error "Failed to encode request; Validation failed for request"
// @Failure 401 {object} response.Error "Invalid old password"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /accounts/change-password [patch]
func (h *Handler) ChangePassword(c *gin.Context) {
	log.Debug().Msg("Account registration")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	accountIDHeader := c.GetHeader("X-Account-ID")
	accountID, err := strconv.Atoi(accountIDHeader)
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
	log.Debug().Int("accountId", accountID).Msg("Header X-Account-ID read successfully")

	var request changePasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Failed to encode request")
		messageID := "FailedToEncodeRequest"
		message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
		if errLoc != nil {
			message = h.EngLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
		}
		c.JSON(http.StatusBadRequest, response.Error{
			Message: message,
			Reason:  err.Error(),
		})
		return
	}
	log.Debug().Msg("Request encoded successfully")

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		log.Error().Err(err).Msg("Validation failed for request")
		messageID := "ValidationFailedForRequest"
		message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
		if errLoc != nil {
			message = h.EngLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
		}
		c.JSON(http.StatusBadRequest, response.Error{
			Message: message,
			Reason:  err.Error(),
		})
		return
	}

	err = h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = h.AccountService.ChangePassword(tx, accountID, request.OldPassword, request.NewPassword)
		if err != nil {
			return err
		}
		err = h.TokenService.DeleteByAccount(tx, accountID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to register account")
		if _, ok := err.(errors.Unauthorized); ok {
			messageID := "InvalidOldPassword"
			message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
			if errLoc != nil {
				message = h.EngLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
			}
			c.JSON(http.StatusConflict, response.Error{
				Message: message,
				Reason:  err.Error(),
			})
			return
		} else {
			messageID := "FailedToRegisterAccount"
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

	log.Debug().Msg("Password changed")
	c.Status(http.StatusNoContent)
}
