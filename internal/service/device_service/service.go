package device_service

import "wakarimi-authentication/internal/model/device"

type Service struct {
	deviceRepo device.Repository
}

func New(deviceRepo device.Repository) *Service {
	return &Service{
		deviceRepo: deviceRepo,
	}
}
