package use_case

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
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
	err = u.refreshTokenService.DeleteByDevice(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to delete device's refresh tokens")
		return err
	}

	return nil
}
