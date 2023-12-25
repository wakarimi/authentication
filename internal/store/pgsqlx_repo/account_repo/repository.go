package account_repo

import "github.com/jmoiron/sqlx"

type Repository struct {
}

func NewRepository(db *sqlx.DB) (repository *Repository) {
	return &Repository{}
}
