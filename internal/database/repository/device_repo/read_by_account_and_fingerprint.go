package device_repo

import (
	"authentication/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) ReadByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (device model.Device, err error) {
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
		return model.Device{}, err
	}
	err = stmt.Get(&device, args)
	if err != nil {
		log.Error().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Failed to read device")
		return model.Device{}, err
	}

	log.Debug().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Device read successfully")
	return device, nil
}
