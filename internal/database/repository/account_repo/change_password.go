package account_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) ChangePassword(tx *sqlx.Tx, accountID int, hashedPassword string) (err error) {
	log.Debug().Msg("Changing password")

	query := `
        UPDATE accounts
        SET hashed_password = :hashed_password
        WHERE id = :account_id
    `

	params := map[string]interface{}{
		"hashed_password": hashedPassword,
		"account_id":      accountID,
	}

	result, err := tx.NamedExec(query, params)
	if err != nil {
		log.Error().Err(err).Msg("Failed to change password")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get rows affected")
		return err
	}

	if rowsAffected == 0 {
		err := fmt.Errorf("no account found with id: %d", accountID)
		log.Error().Err(err).Msg("No account found to change password")
		return err
	}

	log.Debug().Msg("Password changed successfully")
	return nil
}
