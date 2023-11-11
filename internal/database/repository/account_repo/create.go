package account_repo

import (
	"authentication/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) Create(tx *sqlx.Tx, account model.Account) (accountID int, err error) {
	log.Debug().Msg("Creating account")

	query := `
		INSERT INTO accounts(username, hashed_password, created_at, last_login)
		VALUES (:username, :hashed_password, :created_at, :last_login)
		RETURNING id
	`
	rows, err := tx.NamedQuery(query, account)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create account")
		return 0, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close rows")
		}
	}(rows)

	if rows.Next() {
		if err := rows.Scan(&accountID); err != nil {
			log.Error().Err(err).Msg("Failed to scan id into filed")
			return 0, err
		}
	} else {
		err := fmt.Errorf("no id returned after account insert")
		log.Error().Err(err).Msg("No id returned after account insert")
		return 0, err
	}

	log.Debug().Int("accountID", accountID).Msg("Account created")
	return accountID, nil
}
