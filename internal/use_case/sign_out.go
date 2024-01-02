package use_case

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/errors"
)

func (u UseCase) SignOut(deviceID int) (err error) {
	log.Debug().Int("deviceId", deviceID).Msg("Sign out by device")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = u.signOut(tx, deviceID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to sign out")
		return err
	}

	return nil
}

func (u UseCase) signOut(tx *sqlx.Tx, deviceID int) (err error) {
	isDeviceExists, err := u.deviceService.IsExists(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to check device existence")
		return err
	}
	if !isDeviceExists {
		err = errors.NotFound{EntityName: fmt.Sprintf("device with id=%d", deviceID)}
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Device not found")
		return err
	}
	err = u.refreshTokenService.DeleteByDevice(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to delete device's refresh tokens")
		return err
	}

	return nil
}
