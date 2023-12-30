package device_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/device"
)

func (s Service) GetByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (device.Device, error) {
	log.Debug().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Getting device by account")

	readDevice, err := s.deviceRepo.ReadByAccountAndFingerprint(tx, accountID, fingerprint)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Failed to read account by username")
		return device.Device{}, err
	}

	log.Debug().Int("accountId", accountID).Str("fingerprint", fingerprint).Int("deviceId", readDevice.ID).Msg("Device got successfully")
	return readDevice, nil
}
