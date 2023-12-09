package token_service

import (
	"authentication/internal/errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) IsExistsByDevice(tx *sqlx.Tx, deviceID int) (exists bool, err error) {
	log.Debug().Int("deviceId", deviceID).Msg("Checking token for device existence")

	isDeviceExists, err := s.DeviceService.IsExists(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to check device existence")
		return false, err
	}
	if !isDeviceExists {
		err := errors.NotFound{Resource: fmt.Sprintf("device with id=%d", deviceID)}
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Device not found")
		return false, err
	}

	exists, err = s.RefreshTokenRepo.IsExistsByDevice(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to check token existence by device")
		return false, err
	}

	log.Debug().Int("deviceId", deviceID).Bool("exists", exists).Msg("Token for device existence checked")
	return exists, nil
}
