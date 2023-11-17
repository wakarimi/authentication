package account_handler

import (
	"authentication/internal/errors"
	"authentication/internal/handler/response"
	"authentication/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
)

type getMeResponse struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

func (h *Handler) GetMe(c *gin.Context) {
	log.Debug().Msg("Getting requester's account")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(h.Bundle, lang)

	accountIDHeader := c.GetHeader("X-Account-ID")
	accountID, err := strconv.Atoi(accountIDHeader)
	if err != nil {
		log.Error().Err(err).Str("accountIDHeader", accountIDHeader).Msg("Invalid X-Account-ID format")
		c.JSON(http.StatusForbidden, response.Error{
			Message: localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "InvalidHeaderFormat",
			}),
			Reason: err.Error(),
		})
		return
	}

	var account model.Account
	var accountRoles []model.AccountRole
	err = h.TransactionManager.WithTransaction(func(tx *sqlx.Tx) (err error) {
		account, err = h.AccountService.Get(tx, accountID)
		if err != nil {
			return err
		}
		accountRoles, err = h.AccountRoleService.GetAllByAccount(tx, account.ID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create account")
		if _, ok := err.(errors.NotFound); ok {
			messageID := "AccountNotFound"
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
			messageID := "FailedToGetAccount"
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

	rolesResponse := make([]string, len(accountRoles))
	for i, accountRole := range accountRoles {
		rolesResponse[i] = string(accountRole.Role)
	}
	log.Debug().Int("accountId", accountID).Msg("Account got")
	c.JSON(http.StatusOK, getMeResponse{
		ID:       account.ID,
		Username: account.Username,
		Roles:    rolesResponse,
	})
}
