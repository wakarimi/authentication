package device_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) Delete(tx *sqlx.Tx, deviceID int) (err error) {
	log.Debug().Int("deviceId", deviceID).Msg("Deleting device")

	query := `
		DELETE FROM devices
		WHERE id = :id
	`
	args := map[string]interface{}{
		"id": deviceID,
	}
	_, err = tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to delete device")
		return err
	}

	log.Debug().Int("deviceId", deviceID).Msg("Device deleted")
	return nil
}
