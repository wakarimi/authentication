package use_case

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/errors"
	"wakarimi-authentication/internal/model/device"
)

func (u UseCase) SignIn(username string, password string, fingerprint string) (refreshToken string, accessToken string, err error) {
	log.Debug().Msg("Sign in to account")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		refreshToken, accessToken, err = u.signIn(tx, username, password, fingerprint)
		if err != nil {
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

func (u UseCase) signIn(tx *sqlx.Tx, username string, password string, fingerprint string) (refreshToken string, accessToken string, err error) {
	if len(username) == 0 || len(password) == 0 {
		err = errors.Conflict{"username or password not specified"}
		log.Error().Err(err).Str("username", username).Msg("username or password not specified")
		return "", "", err
	}
	isUsernameTaken, err := u.accountService.IsUsernameTaken(tx, username)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check account by username existence")
		return "", "", err
	}
	if !isUsernameTaken {
		err := errors.Unauthorized{Message: "invalid username"}
		log.Error().Err(err).Str("username", username).Msg("Invalid username")
		return "", "", err
	}

	account, err := u.accountService.GetByUsername(tx, username)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read account by username")
		return "", "", err
	}
	isPasswordMatch := u.accountService.IsPasswordMatch(password, account.HashedPassword)
	if !isPasswordMatch {
		err := errors.Unauthorized{Message: "invalid password"}
		log.Error().Err(err).Str("username", username).Msg("Invalid username or password")
		return "", "", err
	}

	isDeviceExists, err := u.deviceService.IsExistsByAccountAndFingerprint(tx, account.ID, fingerprint)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check device existence")
		return "", "", err
	}
	if isDeviceExists {
		foundDevice, err := u.deviceService.GetByAccountAndFingerprint(tx, account.ID, fingerprint)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get device")
			return "", "", err
		}
		err = u.refreshTokenService.DeleteByDevice(tx, foundDevice.ID)
		if err != nil {
			log.Error().Err(err).Int("deviceId", foundDevice.ID).Msg("Failed to delete token by device")
			return "", "", err
		}
		err = u.deviceService.Delete(tx, foundDevice.ID)
		if err != nil {
			log.Error().Err(err).Int("deviceId", foundDevice.ID).Msg("Failed to delete device")
			return "", "", err
		}
	}

	deviceToCreate := device.Device{
		AccountID:   account.ID,
		Fingerprint: fingerprint,
	}
	createdDeviceID, err := u.deviceService.Create(tx, deviceToCreate)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create device")
		return "", "", err
	}

	err = u.accountService.UpdateLastSignIn(tx, account.ID)
	if err != nil {
		log.Error().Int("accountId", account.ID).Msg("Failed to update last sign in")
		return "", "", err
	}

	refreshTokenID, err := u.refreshTokenService.GenerateAndCreateInDatabase(tx, account.ID, createdDeviceID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate and save refresh token")
		return "", "", err
	}
	refreshTokenStruct, err := u.refreshTokenService.Get(tx, refreshTokenID)
	if err != nil {
		log.Error().Msg("Failed to get refresh token")
		return "", "", err
	}
	refreshToken = refreshTokenStruct.Token

	roles, err := u.accountRoleService.GetAllByAccount(tx, account.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get account's roles")
		return "", "", err
	}

	accessToken, err = u.accessTokenService.Generate(refreshTokenID, account.ID, createdDeviceID, roles)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate access token")
		return "", "", err
	}

	return refreshToken, accessToken, nil
}
