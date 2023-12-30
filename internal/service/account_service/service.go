package account_service

import (
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/model/account"
)

type accountRepo interface {
	CountAccounts(tx *sqlx.Tx) (int, error)
	Create(tx *sqlx.Tx, account account.Account) (int, error)
	IsUsernameTaken(tx *sqlx.Tx, username string) (bool, error)
}

type Service struct {
	accountRepo accountRepo
}

func New(accountRepo accountRepo) *Service {
	return &Service{
		accountRepo: accountRepo,
	}
}
