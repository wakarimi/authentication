package account_repo

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repo interface {
	Create(tx *sqlx.Tx, account model.Account) (accountID int, err error)
	Read(tx *sqlx.Tx, accountID int) (account model.Account, err error)
	ReadByUsername(tx *sqlx.Tx, username string) (account model.Account, err error)
	CountAccounts(tx *sqlx.Tx) (count int, err error)
	UpdateLastLogIn(tx *sqlx.Tx, accountID int) (err error)
	IsExists(tx *sqlx.Tx, accountID int) (isExists bool, err error)
	IsUsernameTaken(tx *sqlx.Tx, username string) (isTaken bool, err error)
	ChangePassword(tx *sqlx.Tx, accountID int, hashedPassword string) (err error)
}

type Repository struct{}

func NewRepository() Repo {
	return &Repository{}
}
