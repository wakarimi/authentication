package token_service

import (
	"authentication/internal/service/token_service/token_payload"

	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog/log"
)

func (s Service) GetRefreshTokenPayload(refreshToken string) (refreshTokenPayload token_payload.RefreshToken, err error) {
	log.Debug().Msg("Getting payload of refresh token")

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.RefreshSecretKey), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token payload")
		return token_payload.RefreshToken{}, err
	}

	refreshTokenPayload.AccountID = int(claims["accountId"].(float64))
	refreshTokenPayload.DeviceID = int(claims["deviceId"].(float64))
	refreshTokenPayload.IssuedAt = int64(claims["issuedAt"].(float64))
	refreshTokenPayload.ExpiryAt = int64(claims["expiryAt"].(float64))

	log.Debug().Msg("Payload of refresh token got")
	return refreshTokenPayload, nil
}
