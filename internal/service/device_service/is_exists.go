package device_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) IsExists(tx *sqlx.Tx, deviceID int) (bool, error) {
	log.Debug().Int("deviceId", deviceID).Msg("Checking device existence")

	isExists, err := s.deviceRepo.IsExists(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to check if the username is taken")
		return false, err
	}

	log.Debug().Int("deviceId", deviceID).Bool("isExists", isExists).Msg("The usage of the user name has been checked")
	return isExists, nil
}
