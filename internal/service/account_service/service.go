package account_service

import (
	"authentication/internal/database/repository/account_repo"
	"authentication/internal/service/account_role_service"
)

type Service struct {
	AccountRepo        account_repo.Repo
	AccountRoleService account_role_service.Service
}

func NewService(accountRepo account_repo.Repo,
	accountRoleService account_role_service.Service) (s *Service) {

	s = &Service{
		AccountRepo:        accountRepo,
		AccountRoleService: accountRoleService,
	}

	return s
}
