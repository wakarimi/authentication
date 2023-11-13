package token_service

import (
	"authentication/internal/errors"
	"authentication/internal/handler/token_handler/token_payload"
	"authentication/internal/model"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

const (
	RefreshTokenDuration = time.Hour * 24 * 14
)

func (s Service) GenerateRefreshTokenByCredentials(tx *sqlx.Tx, secretKey string, username string, password string, fingerprint string) (refreshToken string, err error) {
	account, err := s.AccountService.GetByUsername(tx, username)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Failed to get account by username")
		return "", err
	}
	isMatch := CheckPasswordHash(password, account.HashedPassword)
	if !isMatch {
		err := errors.Unauthorized{Message: "invalid username or password"}
		log.Error().Err(err).Str("ureaname", username).Msg("Invalid username or password")
		return "", err
	}

	isExistsByAccountAndFingerprint, err := s.DeviceService.IsExistsByAccountAndFingerprint(tx, account.ID, fingerprint)
	if err != nil {
		log.Error().Err(err).Str("ureaname", username).Str("fingerprint", fingerprint).Msg("Failed to check device existence")
		return "", err
	}
	if isExistsByAccountAndFingerprint {
		foundDevice, err := s.DeviceService.GetByAccountAndFingerprint(tx, account.ID, fingerprint)
		if err != nil {
			log.Error().Err(err).Str("ureaname", username).Str("fingerprint", fingerprint).Msg("Failed to get device")
			return "", err
		}
		err = s.DeviceService.Delete(tx, foundDevice.ID)
		if err != nil {
			log.Error().Err(err).Str("ureaname", username).Str("fingerprint", fingerprint).Msg("Failed to delete device")
			return "", err
		}
	}

	deviceToCreate := model.Device{
		AccountID:   account.ID,
		Fingerprint: fingerprint,
	}
	createdDeviceID, err := s.DeviceService.Create(tx, deviceToCreate)
	if err != nil {
		log.Error().Err(err).Str("ureaname", username).Str("fingerprint", fingerprint).Msg("Failed to create device")
		return "", err
	}

	err = s.AccountService.UpdateLastLogIn(tx, account.ID)
	if err != nil {
		log.Error().Err(err).Int("accountId", account.ID).Msg("Failed to update login")
		return "", err
	}

	payload := token_payload.RefreshToken{
		AccountID: account.ID,
		DeviceID:  createdDeviceID,
		IssuedAt:  time.Now().Unix(),
		ExpiryAt:  time.Now().Add(RefreshTokenDuration).Unix(),
	}
	claims := jwt.MapClaims{
		"accountId": payload.AccountID,
		"deviceId":  payload.DeviceID,
		"issuedAt":  payload.IssuedAt,
		"expiryAt":  payload.ExpiryAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func CheckPasswordHash(password string, hash string) (isMatch bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
