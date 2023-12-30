package refresh_token_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (s Service) Create(tx *sqlx.Tx, refreshToken refresh_token.RefreshToken) error {
	log.Debug().Msg("Creating a refresh token")

	err := s.refreshTokenRepo.Create(tx, refreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create refresh token")
		return err
	}

	log.Debug().Msg("Refresh token created")
	return nil
}
