package account_role_repo

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repo interface {
	Create(tx *sqlx.Tx, accountRole model.AccountRole) (err error)
	ReadAllByAccount(tx *sqlx.Tx, accountID int) (accountRoles []model.AccountRole, err error)
	HasRole(tx *sqlx.Tx, accountID int, roleName model.RoleName) (hasRole bool, err error)
}

type Repository struct{}

func NewRepository() Repo {
	return &Repository{}
}
