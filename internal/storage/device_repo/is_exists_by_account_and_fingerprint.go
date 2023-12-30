package device_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) IsExistsByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (isExists bool, err error) {
	log.Debug().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Checking device existence")

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM devices
			WHERE account_id = :account_id
			  AND fingerprint = :fingerprint
		)
	`
	args := map[string]interface{}{
		"account_id":  accountID,
		"fingerprint": fingerprint,
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Failed to prepare query")
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
		log.Error().Err(err).Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Failed to check device existence")
		return false, err
	}

	log.Debug().Int("accountId", accountID).Str("fingerprint", fingerprint).Bool("isExists", isExists).Msg("Device existence checked")
	return isExists, nil
}
