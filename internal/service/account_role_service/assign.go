package account_role_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account_role"
)

func (s Service) Assign(tx *sqlx.Tx, accountID int, accountRoleName account_role.RoleName) error {
	log.Debug().Int("accountId", accountID).Str("roleName", string(accountRoleName)).Msg("Assigning role to account")

	accountRole := account_role.AccountRole{
		AccountID: accountID,
		Role:      accountRoleName,
	}
	err := s.accountRoleRepo.Create(tx, accountRole)
	if err != nil {
		log.Error().Err(err).Msg("Failed to assign role")
		return err
	}

	log.Debug().Int("accountId", accountID).Str("roleName", string(accountRoleName)).Msg("Role assigned")
	return nil
}
