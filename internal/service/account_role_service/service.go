package account_role_service

import "authentication/internal/database/repository/account_role_repo"

type Service struct {
	AccountRoleRepo account_role_repo.Repo
}

func NewService(accountRoleRepo account_role_repo.Repo) (s *Service) {

	s = &Service{
		AccountRoleRepo: accountRoleRepo,
	}

	return s
}
