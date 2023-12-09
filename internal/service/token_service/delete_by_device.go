package token_service

import (
	"authentication/internal/errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) DeleteByDevice(tx *sqlx.Tx, deviceID int) (err error) {
	log.Debug().Int("deviceId", deviceID).Msg("Deleting refresh token for device")

	isDeviceExists, err := s.DeviceService.IsExists(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to check device existence")
		return err
	}
	if !isDeviceExists {
		err := errors.NotFound{Resource: fmt.Sprintf("device with id=%d", deviceID)}
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Device not found")
		return err
	}

	isExistsByDevice, err := s.IsExistsByDevice(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to check token existence by device")
		return err
	}
	if !isExistsByDevice {
		err := errors.NotFound{Resource: fmt.Sprintf("tokens for device with id=%d", deviceID)}
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Tokens for device not found")
		return err
	}

	err = s.RefreshTokenRepo.DeleteByDevice(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to delete token for device")
		return err
	}

	log.Debug().Int("deviceId", deviceID).Msg("Token by device deleted")
	return nil
}
