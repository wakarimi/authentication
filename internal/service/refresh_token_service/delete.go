package refresh_token_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) Delete(tx *sqlx.Tx, refreshTokenID int) error {
	log.Debug().Int("refreshTokenId", refreshTokenID).Msg("Deleting refresh token")

	err := s.refreshTokenRepo.Delete(tx, refreshTokenID)
	if err != nil {
		log.Error().Err(err).Int("refreshTokenId", refreshTokenID).Msg("Failed to delete refresh token")
		return err
	}

	log.Debug().Int("refreshTokenId", refreshTokenID).Msg("Refresh token deleted")
	return nil
}
