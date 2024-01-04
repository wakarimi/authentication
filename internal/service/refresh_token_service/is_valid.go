package refresh_token_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"time"
)

func (s Service) IsValid(tx *sqlx.Tx, refreshToken string) bool {
	log.Debug().Msg("Checking the refresh token")

	refreshTokenFromDatabase, err := s.refreshTokenRepo.ReadByToken(tx, refreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check token")
		return false
	}

	return refreshTokenFromDatabase.ExpiresAt.After(time.Now())
}
