package account_role_repo

import (
	"authentication/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r Repository) HasRole(tx *sqlx.Tx, accountID int, roleName model.RoleName) (hasRole bool, err error) {
	log.Debug().Int("accountID", accountID).Str("roleName", string(roleName)).Msg("Checking if account already has the role")

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM account_roles
			WHERE account_id = :account_id 
			  AND role = :role
		)
	`
	args := map[string]interface{}{
		"account_id": accountID,
		"role":       string(roleName),
	}

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		log.Error().Err(err).Int("accountID", accountID).Str("roleName", string(roleName)).Msg("Failed to prepare query")
		return false, err
	}
	defer func(stmt *sqlx.NamedStmt) {
		err := stmt.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close statement")
		}
	}(stmt)

	err = stmt.Get(&hasRole, args)
	if err != nil {
		log.Error().Err(err).Int("accountID", accountID).Str("roleName", string(roleName)).Msg("Failed to check if account has the role")
		return false, err
	}

	log.Debug().Int("accountID", accountID).Str("roleName", string(roleName)).Bool("hasRole", hasRole).Msg("Role existence for account checked")
	return hasRole, nil
}
