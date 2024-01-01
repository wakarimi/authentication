package access_token_service

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog/log"
	"time"
	"wakarimi-authentication/internal/model/access_token"
	"wakarimi-authentication/internal/model/account_role"
)

func (s Service) Generate(refreshTokenID int, accountID int, deviceID int, roles []account_role.AccountRole) (string, error) {
	log.Debug().Msg("Generating access token")

	rolesAsString := make([]string, len(roles))
	for i, role := range roles {
		rolesAsString[i] = string(role.Role)
	}

	claims := jwt.MapClaims{
		"accountId":      accountID,
		"deviceId":       deviceID,
		"roles":          rolesAsString,
		"refreshTokenId": refreshTokenID,
		"issuedAt":       time.Now().Unix(),
		"expiryAt":       time.Now().Add(access_token.Duration).Unix(),
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
