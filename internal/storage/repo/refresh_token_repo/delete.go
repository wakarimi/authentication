package refresh_token_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) Delete(tx *sqlx.Tx, refreshTokenID int) (err error) {
	log.Error().Err(err).Msg("Deleting refresh token")

	query := `
		DELETE FROM refresh_tokens
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id": refreshTokenID,
	}
	_, err = tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete refresh token")
		return err
	}

	log.Debug().Msg("Refresh token deleted successfully")
	return nil
}
