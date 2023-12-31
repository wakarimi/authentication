package refresh_token_service

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (s Service) GetPayload(token string) (refreshTokenPayload refresh_token.Payload, err error) {
	log.Debug().Msg("Getting payload of refresh token")

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token payload")
		return refresh_token.Payload{}, err
	}

	refreshTokenPayload.AccountID = int(claims["accountId"].(float64))
	refreshTokenPayload.DeviceID = int(claims["deviceId"].(float64))
	refreshTokenPayload.IssuedAt = int64(claims["issuedAt"].(float64))
	refreshTokenPayload.ExpiryAt = int64(claims["expiryAt"].(float64))

	log.Debug().Msg("Payload of refresh token got")
	return refreshTokenPayload, nil
}
