package token_handler

import (
	"authentication/internal/handler/response"
	"authentication/internal/service/token_service/token_payload"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
)

// verifyRequest represents the request format for verifying access token
type verifyRequest struct {
	// Access token for verifying
	AccessToken string `json:"accessToken" validate:"required"`
}

// verifyResponse represents the reponse format for verifying access token
type verifyResponse struct {
	// Token verification result
	Valid bool `json:"valid"`
	// ID of the account that the token belongs to
	AccountID *int `json:"accountId,omitempty"`
	// ID of the device that the token belongs to
	DeviceID *int `json:"deviceId,omitempty"`
	// Account roles
	Roles *[]string `json:"roles,omitempty"`
	// Time of token issuance
	IssuedAt *int64 `json:"issuedAt,omitempty"`
	// Token expiration time
	ExpiryAt *int64 `json:"expiryAt,omitempty"`
}

// Verify the token
// @Summary Verify the token
// @Tags Tokens
// @Accept json
// @Produce json
// @Param Produce-Language header string false "Language preference" default(en-US)
// @Param request body verifyRequest true "The token to be validated"
// @Success 201 {object} verifyResponse
// @Failure 400 {object} response.Error "Failed to encode request; Valitadion failed for request"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /tokens/verify [post]
func (h *Handler) Verify(c *gin.Context) {
	log.Debug().Msg("Verifying an access token")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	var request verifyRequest
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

	var isValid bool
	var accessTokenPayload token_payload.AccessToken
	err := h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		isValid, err = h.TokenService.IsAccessTokenValid(request.AccessToken)
		if err != nil {
			return err
		}
		if isValid {
			accessTokenPayload, err = h.TokenService.GetAccessTokenPayload(request.AccessToken)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate tokens")
		messageID := "FailedToVerifyToken"
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

	log.Debug().Msg("Tokens validated")
	if isValid {
		c.JSON(http.StatusOK, verifyResponse{
			Valid:     isValid,
			AccountID: &accessTokenPayload.AccountID,
			DeviceID:  &accessTokenPayload.DeviceID,
			Roles:     &accessTokenPayload.Roles,
			IssuedAt:  &accessTokenPayload.IssuedAt,
			ExpiryAt:  &accessTokenPayload.ExpiryAt,
		})
	} else {
		c.JSON(http.StatusOK, verifyResponse{
			Valid: isValid,
		})
	}
}
