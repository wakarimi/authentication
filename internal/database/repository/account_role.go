package repository

import (
	"authentication/internal/models"
	"github.com/jmoiron/sqlx"
)

func AssignRoleTx(tx *sqlx.Tx, userId int, role string) error {
	userRole := models.AccountRole{
		AccountId: userId,
		Role:      role,
	}

	query := `
		INSERT INTO account_roles (account_id, role)
		VALUES (:account_id, :role)
	`
	_, err := tx.NamedExec(query, userRole)
	return err
}
