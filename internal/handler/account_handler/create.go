package account_handler

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

// createRequest represents the request body for registration
type createRequest struct {
	// Desired username
	Username string `json:"username" validate:"required,alphanum"`
	// Password for logging into your account
	Password string `json:"password" validate:"required,alphanum"`
}

// Create an account
// @Summary Creates an account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param Produce-Language header string false "Language preference" default(en-US)
// @Param request body createRequest true "Account credentials"
// @Success 201
// @Failure 400 {object} response.Error "Failed to encode request; Validation failed for request"
// @Failure 409 {object} response.Error "Username is already taken"
// @Failure 500 {object} response.Error "Internal server error"
// @Router /register [post]
func (h *Handler) Create(c *gin.Context) {
	log.Debug().Msg("Creating an account")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	var request createRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Failed to encode request")
		messageID := "FailedToEncodeRequest"
		message, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
		if err != nil {
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
		message, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
		if err != nil {
			message = h.EngLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
		}
		c.JSON(http.StatusBadRequest, response.Error{
			Message: message,
			Reason:  err.Error(),
		})
		return
	}

	err := h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = h.AccountService.Create(tx, request.Username, request.Password)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create account")
		if _, ok := err.(errors.Conflict); ok {
			messageID := "UsernameIsAlreadyTaken"
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
			messageID := "FailedToCreateAccount"
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

	log.Debug().Msg("Account created")
	c.Status(http.StatusCreated)
}