package device_repo

import (
	"authentication/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) Create(tx *sqlx.Tx, device model.Device) (deviceID int, err error) {
	log.Debug().Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Creating device")

	query := `
		INSERT INTO devices(account_id, fingerprint)
		VALUES (:account_id, :fingerprint)
		RETURNING id
	`
	rows, err := tx.NamedQuery(query, device)
	if err != nil {
		log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Failed to create device")
		return 0, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Failed to close rows")
		}
	}(rows)

	if rows.Next() {
		if err := rows.Scan(&deviceID); err != nil {
			log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Failed to scan id into filed")
			return 0, err
		}
	} else {
		err := fmt.Errorf("no id returned after device insert")
		log.Error().Err(err).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("No id returned after device insert")
		return 0, err
	}

	log.Debug().Int("deviceID", deviceID).Int("accountId", device.AccountID).Str("fingerprint", device.Fingerprint).Msg("Device created")
	return deviceID, nil
}
