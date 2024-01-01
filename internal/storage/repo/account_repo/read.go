package account_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (r Repository) Read(tx *sqlx.Tx, accountID int) (readAccount account.Account, err error) {
	log.Debug().Int("accountId", accountID).Msg("Reading account")

	query := `
		SELECT *
		FROM accounts
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id": accountID,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("accountId", accountID).Msg("Failed to prepare query")
		return account.Account{}, err
	}
	err = stmt.Get(&readAccount, args)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to read account")
		return account.Account{}, err
	}

	log.Debug().Int("accountId", accountID).Msg("Account read successfully")
	return readAccount, nil
}
