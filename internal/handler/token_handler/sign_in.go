package token_handler

import (
	"authentication/internal/errors"
	"authentication/internal/handler/response"
	"authentication/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
)

// signInRequest represents the request format for creating first device's tokens
type signInRequest struct {
	// Account username
	Username string `json:"username" validate:"required,alphanum"`
	// Account password
	Password string `json:"password" validate:"required,alphanum"`
	// Device name
	DeviceName string `json:"deviceName"`
	// Device fingerprint
	Fingerprint string `json:"fingerprint" validate:"required"`
}

// signInResponse represents the response format for creating first device's tokens
type signInResponse struct {
	// Refresh token for device
	RefreshToken string `json:"refreshToken"`
	// Access token for device
	AccessToken string `json:"accessToken"`
}

// SignIn to account
// @Summary SignIn to account
// @Tags Tokens
// @Accept json
// @Produce json
// @Param Produce-Language header string false "Language preference" default(en-US)
// @Param request body signInRequest true "Account credentials and device fingerprint"
// @Success 201 {object} signInResponse
// @Failure 400 {object} response.Error "Failed to encode request; Validation failed for request"
// @Failure 401 {object} response.Error "Invalid login or password"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	log.Debug().Msg("Logging in to your account")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	var request signInRequest
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
		device := model.Device{
			Name:        request.DeviceName,
			Fingerprint: request.Fingerprint,
		}
		refreshToken, err = h.TokenService.GenerateRefreshTokenByCredentials(tx, request.Username, request.Password, device)
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
			messageID := "FailedToGetAccess"
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
	c.JSON(http.StatusCreated, signInResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	})
}
