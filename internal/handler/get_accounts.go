package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
	"wakarimi-authentication/internal/handler/response"
)

type getAccountsResponseItem struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type getAccountsResponse struct {
	Accounts []getAccountsResponseItem `json:"accounts"`
}

func (h *Handler) GetAccounts(c *gin.Context) {
	log.Debug().Msg("Getting accounts")

	lang := c.MustGet("lang").(string)
	localizer := i18n.NewLocalizer(&h.bundle, lang)

	idsStr := c.Query("ids")
	idsStrArray := strings.Split(idsStr, ",")
	var ids []int
	for _, idStr := range idsStrArray {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error().Err(err).Msg("Failed to convert id to int")
			messageID := "InvalidIdFormat"
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
		ids = append(ids, id)
	}

	accounts, err := h.useCase.GetAccounts(ids)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get")
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

	responseItems := make([]getAccountsResponseItem, len(accounts))
	for i := range responseItems {
		responseItems[i].ID = accounts[i].ID
		responseItems[i].Username = accounts[i].Username
	}
	log.Debug().Msg("Accounts got")
	c.JSON(http.StatusOK, getAccountsResponse{
		Accounts: responseItems,
	})
}
