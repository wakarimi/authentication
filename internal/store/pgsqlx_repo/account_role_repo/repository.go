package account_role_repo

import (
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/store/pgsqlx_repo"
)

type Repository struct {
	pgsqlx_repo.Transactor
}

func NewPgSQLRepository(db *sqlx.DB) (repository *Repository) {
	return &Repository{
		pgsqlx_repo.NewTransactor(db),
	}
}
