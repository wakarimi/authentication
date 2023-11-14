package device_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsExists(tx *sqlx.Tx, deviceID int) (isExists bool, err error) {
	log.Debug().Int("deviceId", deviceID).Msg("Checking device existence")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM devices
			WHERE id = :id
		)
	`
	args := map[string]interface{}{
		"id": deviceID,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("deviceId", deviceID).Msg("Failed to prepare query")
		return false, err
	}
	defer func(stmt *sqlx.NamedStmt) {
		err := stmt.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close statement")
		}
	}(stmt)

	err = stmt.Get(&isExists, args)
	if err != nil {
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to check device existence")
		return false, err
	}

	log.Debug().Int("deviceId", deviceID).Bool("isExists", isExists).Msg("Device existence checked")
	return isExists, nil
}
