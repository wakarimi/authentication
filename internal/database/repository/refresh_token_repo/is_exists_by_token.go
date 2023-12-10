package refresh_token_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsExistsByToken(tx *sqlx.Tx, token string) (isExists bool, err error) {
	log.Debug().Msg("Checking refresh token existence")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM refresh_tokens
			WHERE token = :token
		)
	`
	args := map[string]interface{}{
		"token": token,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Msg("Failed to prepare query")
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
		log.Error().Err(err).Msg("Failed to check refresh token existence")
		return false, err
	}

	log.Debug().Bool("isExists", isExists).Msg("Token existence checked")
	return isExists, nil
}
