package token_handler

import (
	"authentication/internal/errors"
	"authentication/internal/handler/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"net/http"
)

type logInRequest struct {
	Username    string `json:"username" validate:"required,alphanum"`
	Password    string `json:"password" validate:"required,alphanum"`
	Fingerprint string `json:"fingerprint"`
}

type logInResponse struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

func (h *Handler) LogIn(c *gin.Context) {
	log.Debug().Msg("Creating a first tokens for device")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	var request logInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Failed to encode request")
		c.JSON(http.StatusBadRequest, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "FailedToEncodeRequest"}),
			Reason: err.Error(),
		})
		return
	}
	log.Debug().Str("username", request.Username).Msg("Request encoded successfully")

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		log.Error().Err(err).Msg("Validation failed for request")
		c.JSON(http.StatusBadRequest, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "ValidationFailedForRequest"}),
			Reason: err.Error(),
		})
		return
	}

	var refreshToken string
	var accessToken string
	secretKey := "CHANGE_ME"
	err := h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		refreshToken, err = h.TokenService.GenerateRefreshTokenByCredentials(tx, secretKey, request.Username, request.Password, request.Fingerprint)
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
			c.JSON(http.StatusUnauthorized, response.Error{
				Message: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "InvalidUsernameOrPassword"}),
				Reason: err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, response.Error{
				Message: localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID: "FailedToGenerateTokens"}),
				Reason: err.Error(),
			})
			return
		}
	}

	log.Debug().Msg("Tokens generated")
	c.JSON(http.StatusCreated, logInResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	})
}
