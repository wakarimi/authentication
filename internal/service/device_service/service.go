package device_service

import (
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/model/device"
)

type deviceRepo interface {
	Create(tx *sqlx.Tx, deviceToCreate device.Device) (int, error)
	Delete(tx *sqlx.Tx, deviceID int) error
	ReadByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (device.Device, error)
	IsExistsByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (bool, error)
	IsExists(tx *sqlx.Tx, deviceID int) (bool, error)
}

type Service struct {
	deviceRepo deviceRepo
}

func New(deviceRepo deviceRepo) *Service {
	return &Service{
		deviceRepo: deviceRepo,
	}
}
