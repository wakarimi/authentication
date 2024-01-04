package refresh_token_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) DeleteByAccount(tx *sqlx.Tx, accountID int) error {
	log.Debug().Int("accountID", accountID).Msg("Deleting refresh tokens by account")

	err := s.refreshTokenRepo.DeleteByAccount(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountID", accountID).Msg("Failed to delete refresh tokens")
		return err
	}

	log.Debug().Int("accountID", accountID).Msg("Refresh tokens deleted by account")
	return nil
}
