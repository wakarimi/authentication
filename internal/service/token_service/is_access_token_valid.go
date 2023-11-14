package token_service

import "github.com/rs/zerolog/log"

func (s Service) IsAccessTokenValid(accessToken string) (isValid bool, err error) {
	log.Debug().Msg("Checking the validity of the access token")

	parsedAccessToken, err := s.ParseAccessToken(accessToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse refresh token")
		return false, err
	}

	isValid = parsedAccessToken.Valid
	log.Debug().Bool("isValid", isValid).Msg("Access token checked")
	return isValid, nil
}
