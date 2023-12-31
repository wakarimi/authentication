package refresh_token_service

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"time"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (s Service) GenerateAndCreateInDatabase(tx *sqlx.Tx, accountID int, deviceID int) (string, error) {
	log.Debug().Msg("Generating refresh token")

	issuedAt := time.Now().Unix()
	expiryAt := time.Now().Add(refresh_token.Duration).Unix()

	claims := jwt.MapClaims{
		"accountId": accountID,
		"deviceId":  deviceID,
		"issuedAt":  issuedAt,
		"expiryAt":  expiryAt,
	}

	refreshTokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := refreshTokenByte.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create new token string")
		return "", err
	}

	log.Debug().Msg("Creating refresh token")

	refreshTokenForDatabase := refresh_token.RefreshToken{
		DeviceID:  deviceID,
		Token:     refreshToken,
		CreatedAt: time.Unix(issuedAt, 0),
		ExpiresAt: time.Unix(expiryAt, 0),
	}
	err = s.refreshTokenRepo.Create(tx, refreshTokenForDatabase)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create refresh token")
		return "", err
	}

	log.Debug().Msg("Refresh token generated")
	return refreshToken, nil
}
