package refresh_token_service

type refreshTokenRepo interface {
}

type Service struct {
	refreshTokenRepo refreshTokenRepo
}

func New(refreshTokenRepo refreshTokenRepo) *Service {
	return &Service{
		refreshTokenRepo: refreshTokenRepo,
	}
}
