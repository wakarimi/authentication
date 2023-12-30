package refresh_token_service

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (s Service) Generate(payload refresh_token.Payload) (string, error) {
	log.Debug().Msg("Generating refresh token")

	claims := jwt.MapClaims{
		"accountId": payload.AccountID,
		"deviceId":  payload.DeviceID,
		"issuedAt":  payload.IssuedAt,
		"expiryAt":  payload.ExpiryAt,
	}

	accessTokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := accessTokenByte.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create new token string")
		return "", err
	}

	log.Debug().Msg("Refresh token generated")
	return accessToken, nil
}
