package token_service

import (
	"time"

	"github.com/rs/zerolog/log"
)

func (s Service) IsAccessTokenValid(accessToken string) (isValid bool, err error) {
	log.Debug().Msg("Checking the validity of the access token")

	parsedAccessToken, err := s.ParseAccessToken(accessToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse refresh token")
		return false, err
	}

	isParsed := parsedAccessToken.Valid

	isExpired := false
	if isParsed {
		accessTokenPayload, err := s.GetAccessTokenPayload(accessToken)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get access token payload")
			return false, err
		}
		if accessTokenPayload.ExpiryAt < time.Now().Unix() {
			isExpired = true
		}
	}

	isValid = isParsed && !isExpired

	log.Debug().Bool("isParsed", isParsed).Bool("isExpired", isExpired).Msg("Access token checked")
	return isValid, nil
}
