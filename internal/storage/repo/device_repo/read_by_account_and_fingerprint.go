package device_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/device"
)

func (r Repository) ReadByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (readDevice device.Device, err error) {
	log.Debug().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Reading device")

	query := `
		SELECT *
		FROM devices
		WHERE account_id = :account_id
		  AND fingerprint = :fingerprint
	`
	args := map[string]interface{}{
		"account_id":  accountID,
		"fingerprint": fingerprint,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Failed to prepare query")
		return device.Device{}, err
	}
	err = stmt.Get(&readDevice, args)
	if err != nil {
		log.Error().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Failed to read device")
		return device.Device{}, err
	}

	log.Debug().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Device read successfully")
	return readDevice, nil
}
