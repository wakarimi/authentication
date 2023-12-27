package account_service

import "wakarimi-authentication/internal/model/account"

type Service struct {
	accountRepo account.Repository
}

func NewService(accountRepo account.Repository) *Service {
	return &Service{
		accountRepo: accountRepo,
	}
}
