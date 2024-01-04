package refresh_token_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (r Repository) ReadByToken(tx *sqlx.Tx, token string) (refreshToken refresh_token.RefreshToken, err error) {
	log.Debug().Str("refreshTokenString", token).Msg("Reading refresh token")

	query := `
		SELECT *
		FROM refresh_tokens
		WHERE token = :token
	`
	args := map[string]interface{}{
		"token": token,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Str("refreshTokenString", token).Msg("Failed to prepare query")
		return refresh_token.RefreshToken{}, err
	}
	err = stmt.Get(&refreshToken, args)
	if err != nil {
		log.Error().Str("refreshTokenString", token).Msg("Failed to read refresh token")
		return refresh_token.RefreshToken{}, err
	}

	log.Debug().Str("refreshTokenString", token).Msg("Refresh token read successfully")
	return refreshToken, nil
}
