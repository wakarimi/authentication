package token_service

import (
	"authentication/internal/database/repository/refresh_token_repo"
	"authentication/internal/service/account_role_service"
	"authentication/internal/service/account_service"
	"authentication/internal/service/device_service"
)

type Service struct {
	RefreshTokenRepo   refresh_token_repo.Repo
	RefreshSecretKey   string
	AccessSecretKey    string
	AccountService     account_service.Service
	AccountRoleService account_role_service.Service
	DeviceService      device_service.Service
}

func NewService(refreshTokenRepo refresh_token_repo.Repo,
	refreshSecretKey string,
	accessSecretKey string,
	accountService account_service.Service,
	accountRoleService account_role_service.Service,
	deviceService device_service.Service,
) (s *Service) {
	s = &Service{
		RefreshTokenRepo:   refreshTokenRepo,
		RefreshSecretKey:   refreshSecretKey,
		AccessSecretKey:    accessSecretKey,
		AccountService:     accountService,
		AccountRoleService: accountRoleService,
		DeviceService:      deviceService,
	}

	return s
}
