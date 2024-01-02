package use_case

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"time"
	"wakarimi-authentication/internal/model/access_token"
)

func (u UseCase) VerifyAccessToken(accessToken string) (isValid bool, payload access_token.Payload) {
	log.Debug().Msg("Verifying access token")

	_ = u.transactor.WithTransaction(func(tx *sqlx.Tx) error {
		isValid, payload = u.verifyAccessToken(tx, accessToken)
		return nil
	})

	return isValid, payload
}

func (u UseCase) verifyAccessToken(tx *sqlx.Tx, token string) (bool, access_token.Payload) {
	parsedAccessToken, err := u.accessTokenService.Parse(token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse access token")
		return false, access_token.Payload{}
	}
	if !parsedAccessToken.Valid {
		log.Debug().Msg("Invalid token")
		return false, access_token.Payload{}
	}

	payload, err := u.accessTokenService.GetPayload(token)
	if err != nil {
		log.Debug().Msg("Failed to get payload")
		return false, access_token.Payload{}
	}

	parentRefreshToken, err := u.refreshTokenService.Get(tx, payload.RefreshTokenID)
	if err != nil {
		log.Debug().Msg("Failed to get parent refresh token")
		return false, access_token.Payload{}
	}

	isParentRefreshTokenValid := u.refreshTokenService.IsValid(tx, parentRefreshToken.Token)
	if !isParentRefreshTokenValid {
		log.Debug().Msg("Invalid parent refresh token")
		return false, access_token.Payload{}
	}

	if payload.ExpiryAt < time.Now().Unix() {
		log.Debug().Msg("Expired token")
		return false, access_token.Payload{}
	}

	return true, payload
}
