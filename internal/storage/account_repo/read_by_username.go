package account_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (r Repository) ReadByUsername(tx *sqlx.Tx, username string) (readAccount account.Account, err error) {
	log.Debug().Str("username", username).Msg("Reading account")

	query := `
		SELECT *
		FROM accounts
		WHERE username = :username
	`
	args := map[string]interface{}{
		"username": username,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Str("username", username).Msg("Failed to prepare query")
		return account.Account{}, err
	}
	err = stmt.Get(&readAccount, args)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Failed to read account")
		return account.Account{}, err
	}

	log.Debug().Str("username", username).Msg("Account read successfully")
	return readAccount, nil
}
