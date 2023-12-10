package refresh_token_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) DeleteByAccount(tx *sqlx.Tx, accountID int) (err error) {
	log.Debug().Int("accountId", accountID).Msg("Deleting refresh tokens for account")

	query := `
		DELETE FROM refresh_tokens
		WHERE device_id IN (
			SELECT id FROM devices WHERE account_id = :account_id
		)
	`
	args := map[string]interface{}{
		"account_id": accountID,
	}
	_, err = tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to delete refresh tokens for account")
		return err
	}

	log.Debug().Int("accountId", accountID).Msg("Refresh tokens deleted successfully for account")
	return nil
}
