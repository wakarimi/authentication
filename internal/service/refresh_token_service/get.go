package refresh_token_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (s Service) Get(tx *sqlx.Tx, refreshTokenID int) (refresh_token.RefreshToken, error) {
	log.Debug().Msg("Getting refresh token")

	refreshToken, err := s.refreshTokenRepo.Read(tx, refreshTokenID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token")
		return refresh_token.RefreshToken{}, err
	}

	log.Debug().Msg("Refresh token got")
	return refreshToken, nil
}
