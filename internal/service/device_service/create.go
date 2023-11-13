package device_service

import (
	"authentication/internal/errors"
	"authentication/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) Create(tx *sqlx.Tx, device model.Device) (deviceID int, err error) {
	log.Debug().Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Creating device")

	isAccountExists, err := s.AccountService.IsExists(tx, device.AccountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Failed to check account existence")
		return 0, err
	}
	if !isAccountExists {
		err = errors.NotFound{Resource: fmt.Sprintf("account with id=%d", device.AccountID)}
		log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Account not fount")
		return 0, err
	}

	isDeviceExists, err := s.IsExistsByAccountAndFingerprint(tx, device.AccountID, device.Fingerprint)
	if err != nil {
		log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Failed to check device existence")
		return 0, err
	}
	if isDeviceExists {
		existsDevice, err := s.GetByAccountAndFingerprint(tx, device.AccountID, device.Fingerprint)
		if err != nil {
			log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Failed to get device")
			return 0, err
		}
		err = s.Delete(tx, existsDevice.ID)
		if err != nil {
			log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Failed to delete device")
			return 0, err
		}
	}

	deviceID, err = s.DeviceRepo.Create(tx, device)
	if err != nil {
		log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Failed to create device")
		return 0, err
	}

	log.Debug().Int("deviceId", deviceID).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Device created")
	return deviceID, nil
}
