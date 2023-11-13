package device_service

import (
	"github.com/jmoiron/sqlx"
)

func (s Service) IsExistsByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (isExists bool, err error) {
	isExists, err = s.DeviceRepo.IsExistsByAccountAndFingerprint(tx, accountID, fingerprint)
	if err != nil {
		return false, err
	}

	return isExists, nil
}
