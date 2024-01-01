package account_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) UpdatePassword(tx *sqlx.Tx, accountID int, password string) error {
	log.Debug().Int("accountId", accountID).Msg("Updating password")

	query := `
		UPDATE accounts
		SET hashed_password = :password
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id":       accountID,
		"password": password,
	}

	result, err := tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to update password")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to get rows affected after password update")
		return err
	}
	if rowsAffected == 0 {
		err := fmt.Errorf("no rows affected while updating room name")
		log.Error().Err(err).Int("accountId", accountID).Msg("No rows affected while updating password")
		return err
	}

	log.Debug().Int("accountId", accountID).Msg("Password updated")
	return err
}
