package account_repo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (r *Repository) Create(ctx context.Context, account account.Account) (accountID int, err error) {
	log.Debug().Msg("Creating account")

	query := `
		INSERT INTO accounts(username, hashed_password, created_at)
		VALUES (:username, :hashed_password, CURRENT_TIMESTAMP)
		RETURNING id
	`
	var rows *sqlx.Rows
	err = r.InTransaction(ctx, func(ctx context.Context) (err error) {
		connection := r.GetConnection(ctx)
		rows, err = connection.NamedQuery(query, account)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create account")
			return err
		}
		defer func(rows *sqlx.Rows) {
			err := rows.Close()
			if err != nil {
				log.Error().Err(err).Msg("Failed to close rows")
			}
		}(rows)
		return nil
	})

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
