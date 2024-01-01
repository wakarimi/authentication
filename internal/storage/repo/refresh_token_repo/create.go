package refresh_token_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (r Repository) Create(tx *sqlx.Tx, token refresh_token.RefreshToken) (err error) {
	log.Debug().Msg("Creating refresh token")

	query := `
    INSERT INTO refresh_tokens(device_id, token, created_at, expires_at)
    VALUES (:device_id, :token, :created_at, :expires_at)
	`
	_, err = tx.NamedExec(query, token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create refresh token")
		return err
	}

	log.Debug().Msg("Refresh token created")
	return nil
}
