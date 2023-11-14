package device_service

import (
	"authentication/internal/database/repository/device_repo"
	"authentication/internal/service/account_service"
)

type Service struct {
	DeviceRepo     device_repo.Repo
	AccountService account_service.Service
}

func NewService(deviceRepo device_repo.Repo,
	accountService account_service.Service) (s *Service) {

	s = &Service{
		AccountService: accountService,
		DeviceRepo:     deviceRepo,
	}

	return s
}
