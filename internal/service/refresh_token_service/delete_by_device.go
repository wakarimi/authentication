package refresh_token_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) DeleteByDevice(tx *sqlx.Tx, deviceID int) error {
	log.Debug().Int("deviceId", deviceID).Msg("Deleting refresh token by device")

	err := s.refreshTokenRepo.DeleteByDevice(tx, deviceID)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to delete refresh token")
		return err
	}

	log.Debug().Int("deviceId", deviceID).Msg("Refresh token deleted by device")
	return nil
}
