package device_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) IsExistsByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (bool, error) {
	log.Debug().Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Checking existence device by account and fingerprint")

	isExists, err := s.deviceRepo.IsExistsByAccountAndFingerprint(tx, accountID, fingerprint)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Str("fingerprint", fingerprint).Msg("Failed to checking existence account by account and fingerprint")
		return false, err
	}

	log.Debug().Int("accountId", accountID).Str("fingerprint", fingerprint).Bool("isExists", isExists).Msg("Device existence checked")
	return isExists, nil
}
