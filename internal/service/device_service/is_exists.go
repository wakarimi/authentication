package device_service

import (
	"github.com/jmoiron/sqlx"
)

func (s Service) IsExists(tx *sqlx.Tx, deviceID int) (isExists bool, err error) {
	isExists, err = s.DeviceRepo.IsExists(tx, deviceID)
	if err != nil {
		return false, err
	}

	return isExists, nil
}
