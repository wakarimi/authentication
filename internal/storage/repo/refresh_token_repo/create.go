package refresh_token_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (r Repository) Create(tx *sqlx.Tx, refreshToken refresh_token.RefreshToken) (refreshTokenID int, err error) {
	log.Debug().Msg("Creating refresh token")

	query := `
		INSERT INTO refresh_tokens(device_id, token, created_at, expires_at)
		VALUES (:device_id, :token, :created_at, :expires_at)
		RETURNING id
	`
	rows, err := tx.NamedQuery(query, refreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create refreshToken")
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
		err := fmt.Errorf("no id returned after refreshToken insert")
		log.Error().Err(err).Msg("No id returned after refreshToken insert")
		return 0, err
	}

	log.Debug().Int("refreshTokenId", refreshTokenID).Msg("RefreshToken created")
	return refreshTokenID, nil
}
