package token_service

import (
	"authentication/internal/service/token_service/token_payload"

	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog/log"
)

func (s Service) GetAccessTokenPayload(accessToken string) (accessTokenPayload token_payload.AccessToken, err error) {
	log.Debug().Msg("Getting payload of access token")

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.AccessSecretKey), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get access token payload")
		return token_payload.AccessToken{}, err
	}

	accessTokenPayload.AccountID = int(claims["accountId"].(float64))
	accessTokenPayload.DeviceID = int(claims["deviceId"].(float64))
	accessTokenPayload.Roles = claims["roles"].([]string)
	accessTokenPayload.IssuedAt = int64(claims["issuedAt"].(float64))
	accessTokenPayload.ExpiryAt = int64(claims["expiryAt"].(float64))

	log.Debug().Msg("Payload of access token got")
	return accessTokenPayload, nil
}
