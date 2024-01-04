package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"net/http"
	"wakarimi-authentication/internal/handler/response"
)

type verifyAccessTokenRequest struct {
	AccessToken string `json:"accessToken" validate:"required"`
}

type verifyAccessTokenResponse struct {
	Valid     bool      `json:"valid"`
	AccountID *int      `json:"accountId,omitempty"`
	DeviceID  *int      `json:"deviceId,omitempty"`
	Roles     *[]string `json:"roles,omitempty"`
	IssuedAt  *int64    `json:"issuedAt,omitempty"`
	ExpiryAt  *int64    `json:"expiryAt,omitempty"`
}

func (h *Handler) VerifyAccessToken(c *gin.Context) {
	log.Debug().Msg("Validating an access token")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(&h.bundle, lang)

	var request verifyAccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Failed to encode request")
		messageID := "FailedToEncodeRequest"
		message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
		if errLoc != nil {
			message = h.engLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
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
			message = h.engLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
		}
		c.JSON(http.StatusBadRequest, response.Error{
			Message: message,
			Reason:  err.Error(),
		})
		return
	}

	isValid, payload := h.useCase.VerifyAccessToken(request.AccessToken)

	log.Debug().Msg("Tokens validated")
	if isValid {
		c.JSON(http.StatusOK, verifyAccessTokenResponse{
			Valid:     isValid,
			AccountID: &payload.AccountID,
			DeviceID:  &payload.DeviceID,
			Roles:     &payload.Roles,
			IssuedAt:  &payload.IssuedAt,
			ExpiryAt:  &payload.ExpiryAt,
		})
	} else {
		c.JSON(http.StatusOK, verifyAccessTokenResponse{
			Valid: isValid,
		})
	}
}
