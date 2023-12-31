package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"net/http"
	"wakarimi-authentication/internal/errors"
	"wakarimi-authentication/internal/handler/response"
)

type refreshTokensRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type refreshTokensResponse struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

func (h *Handler) RefreshTokens(c *gin.Context) {
	log.Debug().Msg("Refreshing tokens")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(&h.bundle, lang)

	var request refreshTokensRequest
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

	refreshToken, accessToken, err := h.useCase.RefreshTokens(request.RefreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate tokens")
		if _, ok := err.(errors.Unauthorized); ok {
			messageID := "InvalidPreviousToken"
			message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
			if errLoc != nil {
				message = h.engLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
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
				message = h.engLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
			}
			c.JSON(http.StatusInternalServerError, response.Error{
				Message: message,
				Reason:  err.Error(),
			})
			return
		}
	}

	log.Debug().Msg("Tokens refreshed")
	c.JSON(http.StatusCreated, refreshTokensResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	})
}
