package refresh_token_repo

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repo interface {
	Create(tx *sqlx.Tx, refreshToken model.RefreshToken) (refreshTokenID int, err error)
	ReadByToken(tx *sqlx.Tx, token string) (refreshToken model.RefreshToken, err error)
	DeleteByDevice(tx *sqlx.Tx, deviceID int) (err error)
	IsExists(tx *sqlx.Tx, refreshTokenID int) (exists bool, err error)
	IsExistsByToken(tx *sqlx.Tx, token string) (exists bool, err error)
	IsExistsByDevice(tx *sqlx.Tx, deviceID int) (exists bool, err error)
}

type Repository struct{}

func NewRepository() Repo {
	return &Repository{}
}
