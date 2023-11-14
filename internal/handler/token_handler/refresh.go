package token_handler

import (
	"authentication/internal/errors"
	"authentication/internal/handler/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
)

// refreshRequest represents the request format for refreshing tokens
type refreshRequest struct {
	// Previous refsesh token
	RefreshToken string `json:"refreshToken" validate:"required"`
}

// refreshResponse represents the response format for refreshing tokens
type refreshResponse struct {
	// Refresh token for device
	RefreshToken string `json:"refreshToken"`
	// Access token for device
	AccessToken string `json:"accessToken"`
}

// Refresh tokens
// @Summary Refresh tokens
// @Tags Tokens
// @Accept json
// @Produce json
// @Param Produce-Language header string false "Language preference" default(en-US)
// @Param request body refreshRequest true "Previous refresh token"
// @Success 201 {object} refreshResponse
// @Failure 400 {object} response.Error "Failed to encode request; Valitadion failed for request"
// @Failure 401 {object} response.Error "Invalid token; Expired token"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /tokens/refresh [post]
func (h *Handler) Refresh(c *gin.Context) {
	log.Debug().Msg("Refreshing tokens")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	var request refreshRequest
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

	var refreshToken string
	var accessToken string
	err := h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		refreshToken, err = h.TokenService.GenerateRefreshTokenByRefreshToken(tx, request.RefreshToken)
		if err != nil {
			return err
		}
		accessToken, err = h.TokenService.GenerateAccessToken(tx, refreshToken)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate tokens")
		if _, ok := err.(errors.Unauthorized); ok {
			messageID := "InvalidPreviousToken"
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
			messageID := "FailedToGenerateTokens"
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

	log.Debug().Msg("Tokens refreshed")
	c.JSON(http.StatusCreated, refreshResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	})
}
