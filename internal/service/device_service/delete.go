package device_service

import (
	"authentication/internal/errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (s Service) Delete(tx *sqlx.Tx, deviceID int) (err error) {
	isExists, err := s.IsExists(tx, deviceID)
	if err != nil {
		return err
	}
	if !isExists {
		err = errors.NotFound{Resource: fmt.Sprintf("device with id=%d", deviceID)}
		return err
	}

	err = s.DeviceRepo.Delete(tx, deviceID)
	if err != nil {
		return err
	}

	return nil
}
