package refresh_token_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (r Repository) Read(tx *sqlx.Tx, refreshTokenID int) (refreshToken refresh_token.RefreshToken, err error) {
	log.Debug().Int("refreshTokenId", refreshTokenID).Msg("Reading refresh token")

	query := `
		SELECT *
		FROM refresh_tokens
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id": refreshTokenID,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("refreshTokenId", refreshTokenID).Msg("Failed to prepare query")
		return refresh_token.RefreshToken{}, err
	}
	err = stmt.Get(&refreshToken, args)
	if err != nil {
		log.Error().Int("refreshTokenId", refreshTokenID).Msg("Failed to read refresh token")
		return refresh_token.RefreshToken{}, err
	}

	log.Debug().Int("refreshTokenId", refreshTokenID).Msg("Refresh token read successfully")
	return refreshToken, nil
}
