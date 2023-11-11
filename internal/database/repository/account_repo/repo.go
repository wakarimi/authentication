package account_repo

import (
	"authentication/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repo interface {
	Create(tx *sqlx.Tx, account model.Account) (accountID int, err error)
	CountAccounts(tx *sqlx.Tx) (count int, err error)
	IsUsernameTaken(tx *sqlx.Tx, username string) (isTaken bool, err error)
}

type Repository struct {
}

func NewRepository() Repo {
	return &Repository{}
}
