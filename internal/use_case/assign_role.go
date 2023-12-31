package use_case

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/errors"
	"wakarimi-authentication/internal/model/account_role"
)

func (u UseCase) AssignRole(requesterID int, accountID int, roleName string) (err error) {
	log.Debug().Msg("Assigning role")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = u.assignRole(tx, requesterID, accountID, roleName)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to assign role")
		return err
	}

	return nil
}

func (u UseCase) assignRole(tx *sqlx.Tx, requesterID int, accountID int, roleName string) (err error) {
	role, err := u.accountRoleService.StringToRole(roleName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert role name to role")
		return err
	}

	requesterIsAdmin, err := u.accountRoleService.HasRole(tx, requesterID, account_role.RoleAdmin)
	if !requesterIsAdmin {
		err := errors.Forbidden{Message: "not enough permission"}
		log.Error().Err(err).Int("requesterId", requesterID).Msg("Not enough permission")
		return err
	}

	alreadyHasRole, err := u.accountRoleService.HasRole(tx, accountID, role)
	if alreadyHasRole {
		err := errors.Conflict{Message: "the account already has this role"}
		log.Error().Err(err).Int("accountId", accountID).Str("roleName", roleName).Msg("The account already has this role")
		return err
	}

	err = u.accountRoleService.Assign(tx, accountID, role)
	if err != nil {
		log.Error().Err(err).Msg("Failed to assign role")
		return err
	}

	return nil
}
