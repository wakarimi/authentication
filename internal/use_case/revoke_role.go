package use_case

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/errors"
	"wakarimi-authentication/internal/model/account_role"
)

func (u UseCase) RevokeRole(requesterID int, accountID int, roleName string) (err error) {
	log.Debug().Int("requesterId", requesterID).Int("accountId", accountID).Str("roleName", roleName).Msg("Revoking role")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = u.revokeRole(tx, requesterID, accountID, roleName)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to revoke role")
		return err
	}

	return nil
}

func (u UseCase) revokeRole(tx *sqlx.Tx, requesterID int, accountID int, roleName string) (err error) {
	isRequesterExists, err := u.accountService.IsExists(tx, requesterID)
	if err != nil {
		log.Error().Err(err).Int("requesterId", requesterID).Msg("Failed to check requester's account existence")
		return err
	}
	if !isRequesterExists {
		err = errors.NotFound{EntityName: fmt.Sprintf("account with id=%d", requesterID)}
		log.Error().Err(err).Int("requesterId", requesterID).Msg("Requester's account not found")
		return err
	}

	isAccountExists, err := u.accountService.IsExists(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", requesterID).Msg("Failed to check account existence")
		return err
	}
	if !isAccountExists {
		err = errors.NotFound{EntityName: fmt.Sprintf("account with id=%d", accountID)}
		log.Error().Err(err).Int("accountId", accountID).Msg("Account not found")
		return err
	}

	role, err := u.accountRoleService.StringToRole(roleName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert role name to role")
		return err
	}

	requesterIsAdmin, err := u.accountRoleService.HasRole(tx, requesterID, account_role.RoleAdmin)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check requester's permissions")
		return err
	}
	if !requesterIsAdmin {
		err = errors.Forbidden{"requester is not admin"}
		log.Error().Err(err).Int("requesterId", requesterID).Msg("Requester is not admin")
		return err
	}

	accountHasRole, err := u.accountRoleService.HasRole(tx, accountID, role)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check account's roles")
		return err
	}
	if !accountHasRole {
		err = errors.Conflict{"the account does not have a specified role"}
		log.Error().Err(err).Int("accountId", accountID).Str("roleName", roleName).Msg("The account does not have a specified role")
		return err
	}

	err = u.accountRoleService.Revoke(tx, accountID, role)
	if err != nil {
		log.Error().Err(err).Msg("Failed to revoke role")
		return err
	}

	return err
}
