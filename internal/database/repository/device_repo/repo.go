package device_repo

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repo interface {
	Create(tx *sqlx.Tx, device model.Device) (deviceID int, err error)
	ReadByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (device model.Device, err error)
	Delete(tx *sqlx.Tx, deviceID int) (err error)
	IsExists(tx *sqlx.Tx, deviceID int) (isExists bool, err error)
	IsExistsByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (isExists bool, err error)
}

type Repository struct{}

func NewRepository() Repo {
	return &Repository{}
}
