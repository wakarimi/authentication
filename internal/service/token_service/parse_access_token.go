package token_service

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog/log"
)

func (s Service) ParseAccessToken(accessTokenString string) (accessToken *jwt.Token, err error) {
	log.Debug().Msg("Parsing an access token")

	accessToken, err = jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.RefreshSecretKey), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse refresh token")
		return nil, err
	}

	log.Debug().Msg("Refresh token parsed")
	return accessToken, nil
}
