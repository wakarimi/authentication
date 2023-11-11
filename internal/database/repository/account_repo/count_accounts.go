package account_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) CountAccounts(tx *sqlx.Tx) (count int, err error) {
	log.Debug().Msg("Counting accounts")

	query := `
		SELECT COUNT(*)
		FROM accounts
	`
	err = tx.Get(&count, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to count accounts")
		return 0, err
	}

	log.Debug().Int("count", count).Msg("Account count retrieved")
	return count, nil
}
