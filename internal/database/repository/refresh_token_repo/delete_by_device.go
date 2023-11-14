package refresh_token_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) DeleteByDevice(tx *sqlx.Tx, deviceID int) (err error) {
	log.Error().Err(err).Int("deviceId", deviceID).Msg("Deleting refresh token")

	query := `
		DELETE FROM refresh_tokens
		WHERE device_id = :device_id
	`
	args := map[string]interface{}{
		"device_id": deviceID,
	}
	_, err = tx.NamedExec(query, args)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to delete refresh token")
		return err
	}

	log.Debug().Int("deviceId", deviceID).Msg("Refresh token deleted successfully")
	return nil
}
