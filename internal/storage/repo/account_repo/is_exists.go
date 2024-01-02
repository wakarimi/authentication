package account_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsExists(tx *sqlx.Tx, accountID int) (isExists bool, err error) {
	log.Debug().Int("accountId", accountID).Msg("Checking account existence")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM accounts
			WHERE id = :id
		)
	`
	args := map[string]interface{}{
		"id": accountID,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("accountId", accountID).Msg("Failed to prepare query")
		return false, err
	}
	defer func(stmt *sqlx.NamedStmt) {
		err := stmt.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close statement")
		}
	}(stmt)

	err = stmt.Get(&isExists, args)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to check account existence")
		return false, err
	}

	log.Debug().Int("accountId", accountID).Bool("isExists", isExists).Msg("Account existence checked")
	return isExists, nil
}
