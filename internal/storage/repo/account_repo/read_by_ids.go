package account_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (r Repository) ReadAllByIds(tx *sqlx.Tx, ids []int) ([]account.Account, error) {
	log.Debug().Interface("accountIds", ids).Msg("Reading accounts by IDs")

	query := `
		SELECT *
		FROM accounts
		WHERE id = ANY($1)
	`

	var accounts []account.Account
	err := tx.Select(&accounts, query, pq.Array(ids))
	if err != nil {
		log.Error().Err(err).Interface("accountIds", ids).Msg("Failed to read accounts by IDs")
		return nil, err
	}

	log.Debug().Interface("accountIds", ids).Msg("Accounts read successfully by IDs")
	return accounts, nil
}
