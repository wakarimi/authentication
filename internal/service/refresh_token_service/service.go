package refresh_token_service

import "wakarimi-authentication/internal/model/refresh_token"

type Service struct {
	refreshTokenRepo refresh_token.Repository
}

func NewService(refreshTokenRepo refresh_token.Repository) *Service {
	return &Service{
		refreshTokenRepo: refreshTokenRepo,
	}
}
