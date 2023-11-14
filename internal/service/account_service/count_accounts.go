package account_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) CountAccounts(tx *sqlx.Tx) (count int, err error) {
	log.Debug().Msg("Counting total number of accounts")

	count, err = s.AccountRepo.CountAccounts(tx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to count accounts")
		return 0, err
	}

	log.Debug().Int("count", count).Msg("Total number of accounts retrieved")
	return count, nil
}
