package token_handler

import (
	"authentication/internal/service"
	"authentication/internal/service/token_service"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Handler struct {
	TokenService       token_service.Service
	TransactionManager service.TransactionManager
	Bundle             *i18n.Bundle
	EngLocalizer       *i18n.Localizer
}

func NewHandler(tokenService token_service.Service,
	transactionManager service.TransactionManager,
	bundle *i18n.Bundle,
) (h *Handler) {
	h = &Handler{
		TokenService:       tokenService,
		TransactionManager: transactionManager,
		Bundle:             bundle,
		EngLocalizer:       i18n.NewLocalizer(bundle, "en_US"),
	}

	return h
}
