package use_case

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"time"
	"wakarimi-authentication/internal/errors"
	"wakarimi-authentication/internal/model/access_token"
	"wakarimi-authentication/internal/model/device"
	"wakarimi-authentication/internal/model/refresh_token"
)

func (u UseCase) SignIn(username string, password string, fingerprint string) (refreshToken string, accessToken string, err error) {
	log.Debug().Msg("Sign in to account")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		isUsernameTaken, err := u.accountService.IsUsernameTaken(tx, username)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check account by username existence")
			return err
		}
		if !isUsernameTaken {
			err := errors.Unauthorized{Message: "invalid username"}
			log.Error().Err(err).Str("username", username).Msg("Invalid username")
			return err
		}

		account, err := u.accountService.GetByUsername(tx, username)
		if err != nil {
			log.Error().Err(err).Msg("Failed to read account by username")
			return err
		}
		isPasswordMatch := u.accountService.IsPasswordMatch(password, account.HashedPassword)
		if !isPasswordMatch {
			err := errors.Unauthorized{Message: "invalid password"}
			log.Error().Err(err).Str("username", username).Msg("Invalid username or password")
			return err
		}

		isDeviceExists, err := u.deviceService.IsExistsByAccountAndFingerprint(tx, account.ID, fingerprint)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check device existence")
			return err
		}
		if isDeviceExists {
			foundDevice, err := u.deviceService.GetByAccountAndFingerprint(tx, account.ID, fingerprint)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get device")
				return err
			}
			err = u.refreshTokenService.DeleteByDevice(tx, foundDevice.ID)
			if err != nil {
				log.Error().Err(err).Int("deviceId", foundDevice.ID).Msg("Failed to delete token by device")
				return err
			}
			err = u.deviceService.Delete(tx, foundDevice.ID)
			if err != nil {
				log.Error().Err(err).Int("deviceId", foundDevice.ID).Msg("Failed to delete device")
				return err
			}
		}

		deviceToCreate := device.Device{
			AccountID:   account.ID,
			Fingerprint: fingerprint,
		}
		createdDeviceID, err := u.deviceService.Create(tx, deviceToCreate)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create device")
			return err
		}

		err = u.accountService.UpdateLastSignIn(tx, account.ID)
		if err != nil {
			log.Error().Int("accountId", account.ID).Msg("Failed to update last sign in")
			return err
		}

		refreshTokenPayload := refresh_token.Payload{
			AccountID: account.ID,
			DeviceID:  createdDeviceID,
			IssuedAt:  time.Now().Unix(),
			ExpiryAt:  time.Now().Add(refresh_token.Duration).Unix(),
		}
		refreshToken, err = u.refreshTokenService.Generate(refreshTokenPayload)
		if err != nil {
			log.Error().Err(err).Msg("Failed to generate refresh token")
			return err
		}
		refreshTokenForDatabase := refresh_token.RefreshToken{
			DeviceID:  createdDeviceID,
			Token:     refreshToken,
			CreatedAt: time.Unix(refreshTokenPayload.IssuedAt, 0),
			ExpiresAt: time.Unix(refreshTokenPayload.ExpiryAt, 0),
		}
		err = u.refreshTokenService.Create(tx, refreshTokenForDatabase)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create refresh token in database")
			return err
		}

		accountRoles, err := u.accountRoleService.GetAllByAccount(tx, account.ID)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get account's roles")
			return err
		}
		roles := make([]string, len(accountRoles))
		for i, accountRole := range accountRoles {
			roles[i] = string(accountRole.Role)
		}

		accessTokenPayload := access_token.Payload{
			AccountID: account.ID,
			DeviceID:  createdDeviceID,
			Roles:     roles,
			IssuedAt:  time.Now().Unix(),
			ExpiryAt:  time.Now().Add(access_token.Duration).Unix(),
		}
		accessToken, err = u.accessTokenService.Generate(accessTokenPayload)
		if err != nil {
			log.Error().Err(err).Msg("Failed to generate access token")
			return err
		}

		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to sign in")
		return "", "", err
	}

	return refreshToken, accessToken, nil
}
