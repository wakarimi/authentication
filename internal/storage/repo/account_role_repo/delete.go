package account_role_repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account_role"
)

func (r Repository) Delete(tx *sqlx.Tx, role account_role.AccountRole) error {
	log.Debug().Int("accountID", role.AccountID).Str("roleName", string(role.Role)).Msg("Deleting role from account")

	query := `
		DELETE FROM account_roles
		WHERE account_id = :account_id AND role = :role
	`

	result, err := tx.NamedExec(query, role)
	if err != nil {
		log.Debug().Int("accountID", role.AccountID).Str("roleName", string(role.Role)).Msg("Failed to delete role from account")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get rows affected")
		return err
	}

	if rowsAffected == 0 {
		err = fmt.Errorf("role not found for deletion")
		log.Error().Err(err).Int("accountID", role.AccountID).Str("roleName", string(role.Role)).Msg("Role not found for deletion")
		return err
	}

	log.Debug().Int("accountID", role.AccountID).Str("roleName", string(role.Role)).Msg("Role deleted from account successfully")
	return nil
}
