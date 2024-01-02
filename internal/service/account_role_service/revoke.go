package account_role_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account_role"
)

func (s Service) Revoke(tx *sqlx.Tx, accountID int, accountRoleName account_role.RoleName) error {
	log.Debug().Int("accountId", accountID).Str("roleName", string(accountRoleName)).Msg("Revoking role to account")

	accountRole := account_role.AccountRole{
		AccountID: accountID,
		Role:      accountRoleName,
	}
	err := s.accountRoleRepo.Delete(tx, accountRole)
	if err != nil {
		log.Error().Err(err).Msg("Failed to revoke role")
		return err
	}

	log.Debug().Int("accountId", accountID).Str("roleName", string(accountRoleName)).Msg("Role revoked")
	return nil
}
