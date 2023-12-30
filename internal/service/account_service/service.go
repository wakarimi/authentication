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
}

type Service struct {
	accountRepo accountRepo
}

func New(accountRepo accountRepo) *Service {
	return &Service{
		accountRepo: accountRepo,
	}
}
