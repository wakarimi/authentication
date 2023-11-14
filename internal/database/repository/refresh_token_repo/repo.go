package refresh_token_repo

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repo interface {
	Create(tx *sqlx.Tx, refreshToken model.RefreshToken) (refreshTokenID int, err error)
	ReadByToken(tx *sqlx.Tx, token string) (refreshToken model.RefreshToken, err error)
	Update(tx *sqlx.Tx, refreshTokenID int, refreshToken model.RefreshToken) (err error)
	DeleteByDevice(tx *sqlx.Tx, deviceID int) (err error)
}

type Repository struct{}

func NewRepository() Repo {
	return &Repository{}
}
