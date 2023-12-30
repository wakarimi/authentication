package access_token_service

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/access_token"
)

func (s Service) Generate(payload access_token.Payload) (string, error) {
	log.Debug().Msg("Generating access token")

	claims := jwt.MapClaims{
		"accountId": payload.AccountID,
		"deviceId":  payload.DeviceID,
		"roles":     payload.Roles,
		"issuedAt":  payload.IssuedAt,
		"expiryAt":  payload.ExpiryAt,
	}

	accessTokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := accessTokenByte.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create new token string")
		return "", err
	}

	log.Debug().Msg("Access token generated")
	return accessToken, nil
}
