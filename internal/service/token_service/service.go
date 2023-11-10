package token_service

import (
	"authentication/internal/service/account_role_service"
	"authentication/internal/service/account_service"
	"authentication/internal/service/device_service"
)

type Service struct {
	AccountService     account_service.Service
	AccountRoleService account_role_service.Service
	DeviceService      device_service.Service
}

func NewService(accountService account_service.Service,
	accountRoleService account_role_service.Service,
	deviceService device_service.Service) (s *Service) {

	s = &Service{
		AccountService:     accountService,
		AccountRoleService: accountRoleService,
		DeviceService:      deviceService,
	}

	return s
}
