package account_repo

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) ReadByUsername(tx *sqlx.Tx, username string) (account model.Account, err error) {
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
		return model.Account{}, err
	}
	err = stmt.Get(&account, args)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Failed to read account")
		return model.Account{}, err
	}

	log.Debug().Str("username", username).Msg("Account read successfully")
	return account, nil
}
