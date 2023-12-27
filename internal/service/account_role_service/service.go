package account_role_service

import "wakarimi-authentication/internal/model/account_role"

type Service struct {
	accountRoleRepo account_role.Repository
}

func NewService(accountRoleRepo account_role.Repository) *Service {
	return &Service{
		accountRoleRepo: accountRoleRepo,
	}
}
