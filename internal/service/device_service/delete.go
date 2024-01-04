package device_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) Delete(tx *sqlx.Tx, deviceID int) error {
	log.Debug().Msg("Deleting a device")

	err := s.deviceRepo.Delete(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete device")
		return err
	}

	log.Debug().Msg("Device deleted")
	return nil
}
