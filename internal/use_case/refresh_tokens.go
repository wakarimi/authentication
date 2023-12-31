package use_case

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/errors"
)

func (u UseCase) RefreshTokens(oldRefreshToken string) (refreshToken string, accessToken string, err error) {
	log.Debug().Msg("Refresh tokens")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		refreshToken, accessToken, err = u.refreshTokens(tx, oldRefreshToken)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to refresh token")
		return "", "", err
	}

	return refreshToken, accessToken, err
}

func (u UseCase) refreshTokens(tx *sqlx.Tx, oldRefreshToken string) (refreshToken string, accessToken string, err error) {
	isOldRefreshTokenValid := u.refreshTokenService.IsValid(tx, oldRefreshToken)
	if !isOldRefreshTokenValid {
		err = errors.Unauthorized{Message: "invalid refresh token"}
		log.Error().Err(err).Msg("Invalid token")
		return "", "", err
	}

	oldRefreshTokenFromDatabase, err := u.refreshTokenService.GetByToken(tx, oldRefreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get token from database")
		return "", "", err
	}

	oldRefreshTokenPayload, err := u.refreshTokenService.GetPayload(oldRefreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token's payload")
		return "", "", err
	}

	err = u.refreshTokenService.Delete(tx, oldRefreshTokenFromDatabase.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete old refresh token from database")
		return "", "", err
	}
	refreshToken, err = u.refreshTokenService.GenerateAndCreateInDatabase(tx, oldRefreshTokenPayload.AccountID, oldRefreshTokenPayload.DeviceID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate and save refresh token")
		return "", "", err
	}

	roles, err := u.accountRoleService.GetAllByAccount(tx, oldRefreshTokenPayload.AccountID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get account's roles")
		return "", "", err
	}

	accessToken, err = u.accessTokenService.Generate(oldRefreshTokenPayload.AccountID, oldRefreshTokenPayload.DeviceID, roles)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate access token")
		return "", "", err
	}

	return refreshToken, accessToken, err
}
