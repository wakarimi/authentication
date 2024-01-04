package access_token_service

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/access_token"
)

func (s Service) GetPayload(token string) (payload access_token.Payload, err error) {
	log.Debug().Msg("Getting payload of access token")

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get access token payload")
		return access_token.Payload{}, err
	}

	var rolesSlice []string
	if roles, ok := claims["roles"].([]interface{}); ok {
		for _, role := range roles {
			if strRole, ok := role.(string); ok {
				rolesSlice = append(rolesSlice, strRole)
			} else {
				err := fmt.Errorf("claims is not string slice")
				log.Error().Err(err).Msg("Claim is not []string")
				return access_token.Payload{}, err
			}
		}
	}

	payload.AccountID = int(claims["accountId"].(float64))
	payload.DeviceID = int(claims["deviceId"].(float64))
	payload.Roles = rolesSlice
	payload.RefreshTokenID = int(claims["refreshTokenId"].(float64))
	payload.IssuedAt = int64(claims["issuedAt"].(float64))
	payload.ExpiryAt = int64(claims["expiryAt"].(float64))

	log.Debug().Msg("Payload of access token got")
	return payload, nil
}
