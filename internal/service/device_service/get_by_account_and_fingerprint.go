package device_service

import (
	"authentication/internal/errors"
	"authentication/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (s Service) GetByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (device model.Device, err error) {
	isExists, err := s.IsExistsByAccountAndFingerprint(tx, accountID, fingerprint)
	if err != nil {
		return model.Device{}, err
	}
	if !isExists {
		err = errors.NotFound{Resource: fmt.Sprintf("device with account_id=%d and fingerprint=%s", accountID, fingerprint)}
		return model.Device{}, err
	}

	device, err = s.DeviceRepo.ReadByAccountAndFingerprint(tx, accountID, fingerprint)
	if err != nil {
		return model.Device{}, err
	}

	return device, nil
}
