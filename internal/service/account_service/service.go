package account_service

import (
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/model/account"
)

type accountRepo interface {
	CountAccounts(tx *sqlx.Tx) (int, error)
	Create(tx *sqlx.Tx, accountToCreate account.Account) (int, error)
	IsUsernameTaken(tx *sqlx.Tx, username string) (bool, error)
	ReadByUsername(tx *sqlx.Tx, username string) (account.Account, error)
	UpdateLastSignIn(tx *sqlx.Tx, accountID int) error
	Read(tx *sqlx.Tx, accountID int) (account.Account, error)
	UpdatePassword(tx *sqlx.Tx, accountID int, password string) error
	IsExists(tx *sqlx.Tx, accountID int) (bool, error)
	ReadAllByIds(tx *sqlx.Tx, ids []int) ([]account.Account, error)
}

type Service struct {
	accountRepo accountRepo
}

func New(accountRepo accountRepo) *Service {
	return &Service{
		accountRepo: accountRepo,
	}
}
