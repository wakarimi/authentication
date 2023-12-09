package refresh_token_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (*Repository) IsExistsByDevice(tx *sqlx.Tx, deviceID int) (isExists bool, err error) {
	log.Debug().Int("deviceId", deviceID).Msg("Checking refresh tocken for device existence")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM refresh_tokens
			WHERE device_id = :device_id
		)
	`
	args := map[string]interface{}{
		"device_id": deviceID,
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
		log.Error().Err(err).Int("deviceId", deviceID).Msg("Failed to check account existence")
		return false, err
	}

	log.Debug().Int("deviceId", deviceID).Bool("isExists", isExists).Msg("Token existence by device checked")
	return isExists, nil
}
