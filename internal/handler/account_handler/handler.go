package account_handler

import (
	"authentication/internal/service"
	"authentication/internal/service/account_role_service"
	"authentication/internal/service/account_service"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Handler struct {
	AccountService     account_service.Service
	AccountRoleService account_role_service.Service
	TransactionManager service.TransactionManager
	Bundle             *i18n.Bundle
}

func NewHandler(accountService account_service.Service,
	accountRoleService account_role_service.Service,
	transactionManager service.TransactionManager,
	bundle *i18n.Bundle) (h *Handler) {

	h = &Handler{
		AccountService:     accountService,
		AccountRoleService: accountRoleService,
		TransactionManager: transactionManager,
		Bundle:             bundle,
	}

	return h
}
