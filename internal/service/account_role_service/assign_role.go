package account_role_service

import (
	"authentication/internal/errors"
	"authentication/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) AssignRole(tx *sqlx.Tx, accountID int, roleName model.RoleName) (err error) {
	log.Debug().Int("accountID", accountID).Str("roleName", string(roleName)).Msg("Assigning role to account")

	roleExists, err := s.AccountRoleRepo.HasRole(tx, accountID, roleName)
	if err != nil {
		log.Error().Err(err).Int("accountID", accountID).Str("roleName", string(roleName)).Msg("Failed to check if the account already has the role")
		return err
	}
	if roleExists {
		err = errors.Conflict{Message: fmt.Sprintf("account %d already has role %s", accountID, roleName)}
		log.Error().Err(err).Msg("Account already has the role")
		return err
	}

	accountRoleToCreate := model.AccountRole{
		AccountID: accountID,
		Role:      roleName,
	}
	err = s.AccountRoleRepo.Create(tx, accountRoleToCreate)
	if err != nil {
		log.Error().Err(err).Int("accountID", accountID).Str("roleName", string(roleName)).Msg("Failed to assign role to account")
		return err
	}

	log.Debug().Int("accountID", accountID).Str("roleName", string(roleName)).Msg("Role assigned to account successfully")
	return nil
}
