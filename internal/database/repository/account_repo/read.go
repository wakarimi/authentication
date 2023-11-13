package account_repo

import (
	"authentication/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) Read(tx *sqlx.Tx, accountID int) (account model.Account, err error) {
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
		return model.Account{}, err
	}
	err = stmt.Get(&account, args)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to read account")
		return model.Account{}, err
	}

	log.Debug().Int("accountId", accountID).Msg("Account read successfully")
	return account, nil
}
