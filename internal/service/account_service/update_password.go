package account_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) UpdatePassword(tx *sqlx.Tx, accountID int, hashedNewPassword string) error {
	log.Debug().Int("accountId", accountID).Msg("Updating password")

	err := s.accountRepo.UpdatePassword(tx, accountID, hashedNewPassword)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to update password")
		return err
	}

	log.Debug().Int("accountId", accountID).Msg("Password updated")
	return nil
}
