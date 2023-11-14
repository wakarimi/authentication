package refresh_token_repo

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r *Repository) Update(tx *sqlx.Tx, refreshTokenID int, refreshToken model.RefreshToken) (err error) {
	log.Debug().Int("refreshTokenId", refreshTokenID).Msg("Updating directory")

	query := `
		UPDATE refresh_tokens
    SET device_id = :device_id, token = :token, created_at = :created_at, expires_at = :expires_at
		WHERE id = :id
	`

	refreshToken.ID = refreshTokenID
	_, err = tx.NamedExec(query, refreshToken)

	if err != nil {
		log.Error().Err(err).Int("refreshTokenId", refreshTokenID)
		return err
	}

	log.Debug().Int("refreshTokenId", refreshTokenID).Msg("Directory updated successfully")
	return nil
}
