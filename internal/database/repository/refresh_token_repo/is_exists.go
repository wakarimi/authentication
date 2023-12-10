package refresh_token_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (*Repository) IsExists(tx *sqlx.Tx, refreshTokenID int) (isExists bool, err error) {
	log.Debug().Int("refreshTokenId", refreshTokenID).Msg("Checking refresh token for device existence")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM refresh_tokens
			WHERE id = :id
		)
	`
	args := map[string]interface{}{
		"id": refreshTokenID,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("refreshTokenId", refreshTokenID).Msg("Failed to prepare query")
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
		log.Error().Err(err).Int("refreshTokenId", refreshTokenID).Msg("Failed to check refresh token existence")
		return false, err
	}

	log.Debug().Int("refreshTokenId", refreshTokenID).Bool("isExists", isExists).Msg("Token existence by device checked")
	return isExists, nil
}
