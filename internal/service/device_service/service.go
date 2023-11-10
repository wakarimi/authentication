package device_service

import (
	"authentication/internal/database/repository/device_repo"
)

type Service struct {
	DeviceRepo device_repo.Repo
}

func NewService(deviceRepo device_repo.Repo) (s *Service) {

	s = &Service{
		DeviceRepo: deviceRepo,
	}

	return s
}
