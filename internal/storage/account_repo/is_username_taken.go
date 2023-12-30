package account_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsUsernameTaken(tx *sqlx.Tx, username string) (isTaken bool, err error) {
	log.Debug().Str("username", username).Msg("Checking username availability")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM accounts
			WHERE username = :username
		)
	`
	args := map[string]interface{}{
		"username": username,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Str("username", username).Msg("Failed to prepare query")
		return false, err
	}
	defer func(stmt *sqlx.NamedStmt) {
		err := stmt.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close statement")
		}
	}(stmt)

	err = stmt.Get(&isTaken, args)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Failed to check room existence")
		return false, err
	}

	log.Debug().Str("username", username).Bool("isTaken", isTaken).Msg("Username availability checked")
	return isTaken, nil
}
