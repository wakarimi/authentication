package account_role_service

import (
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/model/account_role"
)

type accountRoleRepo interface {
	Create(tx *sqlx.Tx, role account_role.AccountRole) error
	ReadAllByAccount(tx *sqlx.Tx, accountID int) ([]account_role.AccountRole, error)
	Delete(tx *sqlx.Tx, role account_role.AccountRole) error
}

type Service struct {
	accountRoleRepo accountRoleRepo
}

func New(accountRoleRepo accountRoleRepo) *Service {
	return &Service{
		accountRoleRepo: accountRoleRepo,
	}
}
