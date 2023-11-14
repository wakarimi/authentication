package account_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) UpdateLastLogIn(tx *sqlx.Tx, accountID int) (err error) {
	log.Debug().Int("accountId", accountID).Msg("Updating account's last login")

	query := `
		UPDATE accounts
		SET last_login = CURRENT_TIMESTAMP
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id": accountID,
	}

	result, err := tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to update account's last login")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to get rows affected after account's last login update")
		return err
	}
	if rowsAffected == 0 {
		err := fmt.Errorf("no rows affected while updating room name")
		log.Error().Err(err).Int("accountId", accountID).Msg("No rows affected while updating account's last login")
		return err
	}

	log.Debug().Int("accountId", accountID).Msg("Account's last login name updated")
	return err
}
