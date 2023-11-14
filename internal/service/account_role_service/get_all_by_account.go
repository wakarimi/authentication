package account_role_service

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) GetAllByAccount(tx *sqlx.Tx, accountID int) (accountRoles []model.AccountRole, err error) {
	log.Debug().Int("accountID", accountID).Msg("Getting roles of account")

	accountRoles, err = s.AccountRoleRepo.ReadAllByAccount(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountID", accountID).Msg("Failed to assign role to account")
		return make([]model.AccountRole, 0), err
	}

	log.Debug().Int("accountID", accountID).Msg("Role assigned to account successfully")
	return accountRoles, nil
}
