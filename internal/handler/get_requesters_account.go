package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"wakarimi-authentication/internal/errors"
	"wakarimi-authentication/internal/handler/response"
)

type getRequestersAccountResponse struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

func (h *Handler) GetRequestersAccount(c *gin.Context) {
	log.Debug().Msg("Getting requester's account")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(&h.bundle, lang)

	accountIDHeader := c.GetHeader("X-Account-ID")
	requesterAccountID, err := strconv.Atoi(accountIDHeader)
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

	account, roles, err := h.useCase.GetAccountDetails(requesterAccountID, requesterAccountID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to assign role")
		if _, ok := err.(errors.Forbidden); ok {
			messageID := "NotEnoughPermissions"
			message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
			if errLoc != nil {
				message = h.engLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
			}
			c.JSON(http.StatusUnauthorized, response.Error{
				Message: message,
				Reason:  err.Error(),
			})
			return
		} else if _, ok := err.(errors.NotFound); ok {
			messageID := "AccountNotFound"
			message, errLoc := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
			if errLoc != nil {
				message = h.engLocalizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
			}
			c.JSON(http.StatusConflict, response.Error{
				Message: message,
				Reason:  err.Error(),
			})
			return
		} else {
			messageID := "FailedToGetAccount"
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

	rolesResponse := make([]string, len(roles))
	for i, accountRole := range roles {
		rolesResponse[i] = string(accountRole.Role)
	}
	log.Debug().Int("accountId", requesterAccountID).Msg("Account got")
	c.JSON(http.StatusOK, getRequestersAccountResponse{
		ID:       account.ID,
		Username: account.Username,
		Roles:    rolesResponse,
	})
}
