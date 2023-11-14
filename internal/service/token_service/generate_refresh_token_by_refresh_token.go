package token_service

import (
	"authentication/internal/errors"
	"authentication/internal/model"
	"authentication/internal/service/constants"
	token_payload "authentication/internal/service/token_service/token_payload"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) GenerateRefreshTokenByRefreshToken(tx *sqlx.Tx, oldRefreshToken string) (newRefreshToken string, err error) {
	isValid, err := s.IsRefreshTokenValid(oldRefreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failet to check refresh token")
		return "", err
	}
	if !isValid {
		err = errors.Unauthorized{Message: "invalid token"}
		log.Error().Err(err).Msg("Invalid token")
		return "", err
	}

	oldTokenFromDatabase, err := s.RefreshTokenRepo.ReadByToken(tx, oldRefreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read token")
		return "", err
	}

	oldRefreshTokenPayload, err := s.GetRefreshTokenPayload(oldRefreshToken)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token payload")
		return "", err
	}

	payload := token_payload.RefreshToken{
		AccountID: oldRefreshTokenPayload.AccountID,
		DeviceID:  oldRefreshTokenPayload.DeviceID,
		IssuedAt:  time.Now().Unix(),
		ExpiryAt:  time.Now().Add(constants.RefreshTokenDuration).Unix(),
	}
	newClaims := jwt.MapClaims{
		"accountId": payload.AccountID,
		"deviceId":  payload.DeviceID,
		"issuedAt":  payload.IssuedAt,
		"expiryAt":  payload.ExpiryAt,
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	newRefreshToken, err = newToken.SignedString([]byte(s.RefreshSecretKey))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create new token string")
		return "", err
	}

	tokenForDatabase := model.RefreshToken{
		DeviceID:  payload.DeviceID,
		Token:     newRefreshToken,
		CreatedAt: time.Unix(payload.IssuedAt, 0),
		ExpiresAt: time.Unix(payload.ExpiryAt, 0),
	}

	err = s.RefreshTokenRepo.Update(tx, oldTokenFromDatabase.ID, tokenForDatabase)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update refresh token in database")
		return "", err
	}

	return newRefreshToken, nil
}
