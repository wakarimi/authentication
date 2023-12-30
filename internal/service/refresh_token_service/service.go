package refresh_token_service

import (
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/model/refresh_token"
)

type refreshTokenRepo interface {
	Create(tx *sqlx.Tx, token refresh_token.RefreshToken) error
	DeleteByDevice(tx *sqlx.Tx, deviceID int) error
}

type Service struct {
	secretKey        string
	refreshTokenRepo refreshTokenRepo
}

func New(secretKey string,
	refreshTokenRepo refreshTokenRepo) *Service {
	return &Service{
		secretKey:        secretKey,
		refreshTokenRepo: refreshTokenRepo,
	}
}
