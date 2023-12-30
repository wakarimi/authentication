package account_role_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account_role"
)

func (s Service) Assign(tx *sqlx.Tx, accountRole account_role.AccountRole) error {
	log.Debug().Int("accountId", accountRole.AccountID).Str("roleName", string(accountRole.Role)).Msg("Assigning role to account")

	err := s.accountRoleRepo.Create(tx, accountRole)
	if err != nil {
		log.Error().Err(err).Msg("Failed to assign role")
		return err
	}

	log.Debug().Int("accountId", accountRole.AccountID).Str("roleName", string(accountRole.Role)).Msg("Role assigned")
	return nil
}
