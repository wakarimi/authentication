package account_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) UpdateLastSignIn(tx *sqlx.Tx, accountID int) error {
	log.Debug().Int("accountId", accountID).Msg("Updating last sign in")

	err := s.accountRepo.UpdateLastSignIn(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to update last sing in")
		return err
	}

	log.Debug().Int("accountId", accountID).Msg("Last sign in timestamp updated")
	return nil
}
