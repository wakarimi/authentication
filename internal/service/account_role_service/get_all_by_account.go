package account_role_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account_role"
)

func (s Service) GetAllByAccount(tx *sqlx.Tx, accountID int) ([]account_role.AccountRole, error) {
	log.Debug().Int("accountID", accountID).Msg("Getting roles of account")

	accountRoles, err := s.accountRoleRepo.ReadAllByAccount(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountID", accountID).Msg("Failed to assign role to account")
		return make([]account_role.AccountRole, 0), err
	}

	log.Debug().Int("accountID", accountID).Msg("Role assigned to account successfully")
	return accountRoles, nil
}
