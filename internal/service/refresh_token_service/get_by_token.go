package refresh_token_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (s Service) GetByToken(tx *sqlx.Tx, token string) (refresh_token.RefreshToken, error) {
	log.Debug().Msg("Deleting refresh token")

	refreshToken, err := s.refreshTokenRepo.ReadByToken(tx, token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete refresh token")
		return refresh_token.RefreshToken{}, err
	}

	log.Debug().Msg("Refresh token deleted")
	return refreshToken, nil
}
