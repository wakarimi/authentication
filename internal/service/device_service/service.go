package device_service

type deviceRepo interface {
}

type Service struct {
	deviceRepo deviceRepo
}

func New(deviceRepo deviceRepo) *Service {
	return &Service{
		deviceRepo: deviceRepo,
	}
}
