package account_role_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account_role"
)

func (r Repository) Create(tx *sqlx.Tx, accountRole account_role.AccountRole) (err error) {
	log.Debug().Int("accountID", accountRole.AccountID).Str("roleName", string(accountRole.Role)).Msg("Assigning role to account")

	query := `
		INSERT INTO account_roles(account_id, role)
		VALUES (:account_id, :role)
	`
	rows, err := tx.NamedQuery(query, accountRole)
	if err != nil {
		log.Debug().Int("accountID", accountRole.AccountID).Str("roleName", string(accountRole.Role)).Msg("Failed to assign role")
		return err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close rows")
		}
	}(rows)

	log.Debug().Int("accountID", accountRole.AccountID).Str("roleName", string(accountRole.Role)).Msg("Role assigned to account successfully")
	return nil
}
