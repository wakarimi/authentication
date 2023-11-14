package refresh_token_repo

import (
	"authentication/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) Create(tx *sqlx.Tx, refreshToken model.RefreshToken) (refreshTokenID int, err error) {
	log.Debug().Msg("Creating refresh token")

	query := `
    INSERT INTO refresh_tokens(device_id, token, created_at, expires_at)
    VALUES (:device_id, :token, :created_at, :expires_at)
		RETURNING id
	`
	rows, err := tx.NamedQuery(query, refreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create refresh token")
		return 0, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close rows")
		}
	}(rows)

	if rows.Next() {
		if err := rows.Scan(&refreshTokenID); err != nil {
			log.Error().Err(err).Msg("Failed to scan id into filed")
			return 0, err
		}
	} else {
		err := fmt.Errorf("no id returned after refresh token insert")
		log.Error().Err(err).Msg("No id returned after refresh token insert")
		return 0, err
	}

	log.Debug().Int("refreshTokenId", refreshTokenID).Msg("Refresh token created")
	return refreshTokenID, nil
}
