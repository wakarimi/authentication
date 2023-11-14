package refresh_token_repo

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) ReadByToken(tx *sqlx.Tx, token string) (refreshToken model.RefreshToken, err error) {
	log.Debug().Str("refreshTokenString", token).Msg("Reading device")

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
		return model.RefreshToken{}, err
	}
	err = stmt.Get(&refreshToken, args)
	if err != nil {
		log.Error().Str("refreshTokenString", token).Msg("Failed to read refresh token")
		return model.RefreshToken{}, err
	}

	log.Debug().Str("refreshTokenString", token).Msg("Refresh token read successfully")
	return refreshToken, nil
}
