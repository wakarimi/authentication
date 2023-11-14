package token_service

import "github.com/rs/zerolog/log"

func (s Service) IsRefreshTokenValid(refreshToken string) (isValid bool, err error) {
	log.Debug().Msg("Checking the refresh token")

	parsedRefreshToken, err := s.ParseRefreshToken(refreshToken)
	if err != nil {
		log.Error().Err(err).Str("refreshToken", refreshToken).Msg("Failed to parse refresh token")
		return false, err
	}

	isValid = parsedRefreshToken.Valid
	log.Debug().Bool("isValid", isValid).Msg("Refresh token checked")
	return isValid, nil
}
