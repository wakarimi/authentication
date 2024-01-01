package account_role_repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account_role"
)

func (r Repository) ReadAllByAccount(tx *sqlx.Tx, accountID int) (accountRoles []account_role.AccountRole, err error) {
	log.Debug().Int("accountId", accountID).Msg("Reading account's role by account")

	query := `
		SELECT *
		FROM account_roles
		WHERE account_id = :account_id
	`
	args := map[string]interface{}{
		"account_id": accountID,
	}
	rows, err := tx.NamedQuery(query, args)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to read account's roles")
		return make([]account_role.AccountRole, 0), err
	}
	defer rows.Close()

	for rows.Next() {
		var accountRole account_role.AccountRole
		if err = rows.StructScan(&accountRole); err != nil {
			log.Error().Err(err).Int("accountId", accountID).Msg("Failed to scan song")
			return make([]account_role.AccountRole, 0), err
		}
		accountRoles = append(accountRoles, accountRole)
	}

	log.Debug().Int("accountId", accountID).Int("count", len(accountRoles)).Msg("Account's roles read")
	return accountRoles, nil
}
