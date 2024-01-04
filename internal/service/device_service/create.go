package device_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/device"
)

func (s Service) Create(tx *sqlx.Tx, deviceToCreate device.Device) (int, error) {
	log.Debug().Int("accountId", deviceToCreate.AccountID).Str("fingerprint", deviceToCreate.Fingerprint).Msg("Creating a device")

	accountID, err := s.deviceRepo.Create(tx, deviceToCreate)
	if err != nil {
		log.Error().Err(err).Int("accountId", deviceToCreate.AccountID).Str("fingerprint", deviceToCreate.Fingerprint).Msg("Failed to create device")
		return 0, err
	}

	log.Debug().Int("accountId", deviceToCreate.AccountID).Str("fingerprint", deviceToCreate.Fingerprint).Msg("Device created")
	return accountID, nil
}
