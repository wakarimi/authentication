package account_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) IsExists(tx *sqlx.Tx, accountID int) (isExists bool, err error) {
	log.Debug().Int("accountId", accountID).Msg("Checking account existence")

	isExists, err = s.AccountRepo.IsExists(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to check if the username is taken")
		return false, err
	}

	log.Debug().Int("accountId", accountID).Bool("isExists", isExists).Msg("The usage of the user name has been checked")
	return isExists, nil
}
