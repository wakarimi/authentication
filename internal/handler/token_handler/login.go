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

// loginRequest represents the request format for creating first device's tokens
type loginRequest struct {
	// Account username
	Username string `json:"username" validate:"required,alphanum"`
	// Account password
	Password string `json:"password" validate:"required,alphanum"`
	// Device Fingerprint
	Fingerprint string `json:"fingerprint"`
}

// loginResponse represents the reponse format for creating first device's tokens
type loginResponse struct {
	// Refresh token for device
	RefreshToken string `json:"refreshToken"`
	// Access token for device
	AccessToken string `json:"accessToken"`
}

// Login to account
// @Summary Login to account
// @Tags Tokens
// @Accept json
// @Produce json
// @Param Produce-Language header string false "Language preference" default(en-US)
// @Param request body loginRequest true "Account credentials and device fingerprint"
// @Success 201 {object} loginResponse
// @Failure 400 {object} response.Error "Failed to encode request; Valitadion failed for request"
// @Failure 401 {object} response.Error "Invalid login or password"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	log.Debug().Msg("Loginning to account")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	var request loginRequest
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
	log.Debug().Str("username", request.Username).Msg("Request encoded successfully")

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
		refreshToken, err = h.TokenService.GenerateRefreshTokenByCredentials(tx, request.Username, request.Password, request.Fingerprint)
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
			messageID := "InvalidUsernameOrPassword"
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

	log.Debug().Msg("Tokens generated")
	c.JSON(http.StatusCreated, loginResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	})
}
