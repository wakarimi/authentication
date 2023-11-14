package token_service

import (
	"authentication/internal/errors"
	"authentication/internal/service/constants"
	"authentication/internal/service/token_service/token_payload"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) GenerateAccessToken(tx *sqlx.Tx, refreshToken string) (accessToken string, err error) {
	log.Debug().Msg("Generating access token")

	isRefreshTokenValid, err := s.IsRefreshTokenValid(refreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check refresh token valid")
		return "", err
	}
	if !isRefreshTokenValid {
		err := errors.Unauthorized{Message: "invalid refresh token"}
		log.Error().Err(err).Msg("Invalid refresh token")
		return "", err
	}

	refreshTokenPayload, err := s.GetRefreshTokenPayload(refreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token payload")
		return "", err
	}

	accountRoles, err := s.AccountRoleService.GetAllByAccount(tx, refreshTokenPayload.AccountID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get account's roles")
		return "", err
	}
	roles := make([]string, len(accountRoles))
	for i, accountRole := range accountRoles {
		roles[i] = string(accountRole.Role)
	}

	payload := token_payload.AccessToken{
		AccountID: refreshTokenPayload.AccountID,
		DeviceID:  refreshTokenPayload.AccountID,
		Roles:     roles,
		IssuedAt:  time.Now().Unix(),
		ExpiryAt:  time.Now().Add(constants.AccessTokenDuration).Unix(),
	}
	claims := jwt.MapClaims{
		"accountId": payload.AccountID,
		"deviceId":  payload.DeviceID,
		"roles":     payload.Roles,
		"issuedAt":  payload.IssuedAt,
		"expiryAt":  payload.ExpiryAt,
	}

	accessTokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = accessTokenByte.SignedString([]byte(s.RefreshSecretKey))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create new token string")
		return "", err
	}

	log.Debug().Msg("Access token generated")
	return accessToken, nil
}
