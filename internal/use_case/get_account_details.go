package use_case

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/errors"
	"wakarimi-authentication/internal/model/account"
	"wakarimi-authentication/internal/model/account_role"
)

func (u UseCase) GetAccountDetails(requesterID int, accountID int) (acc account.Account, roles []account_role.AccountRole, err error) {
	log.Debug().Msg("Getting account details")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		acc, roles, err = u.getAccountDetails(tx, requesterID, accountID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get account details")
		return account.Account{}, make([]account_role.AccountRole, 0), err
	}

	return acc, roles, nil
}
func (u UseCase) getAccountDetails(tx *sqlx.Tx, requesterID int, accountID int) (acc account.Account, roles []account_role.AccountRole, err error) {
	isRequesterExists, err := u.accountService.IsExists(tx, requesterID)
	if err != nil {
		log.Error().Err(err).Int("requesterId", requesterID).Msg("Failed to check requester's account existence")
		return account.Account{}, make([]account_role.AccountRole, 0), err
	}
	if !isRequesterExists {
		err = errors.NotFound{EntityName: fmt.Sprintf("account with id=%d", requesterID)}
		log.Error().Err(err).Int("requesterId", requesterID).Msg("Requester's account not found")
		return account.Account{}, make([]account_role.AccountRole, 0), err
	}

	isAccountExists, err := u.accountService.IsExists(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", requesterID).Msg("Failed to check account existence")
		return account.Account{}, make([]account_role.AccountRole, 0), err
	}
	if !isAccountExists {
		err = errors.NotFound{EntityName: fmt.Sprintf("account with id=%d", accountID)}
		log.Error().Err(err).Int("accountId", accountID).Msg("Account not found")
		return account.Account{}, make([]account_role.AccountRole, 0), err
	}

	requesterIsAdmin, err := u.accountRoleService.HasRole(tx, requesterID, account_role.RoleAdmin)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check requester's permissions")
		return account.Account{}, make([]account_role.AccountRole, 0), err
	}

	isAccountOwnedByRequester := requesterID == accountID

	if !requesterIsAdmin && !isAccountOwnedByRequester {
		err = errors.Forbidden{"not enough permission"}
		log.Error().Err(err).Msg("Not enough permission")
		return account.Account{}, make([]account_role.AccountRole, 0), err
	}

	acc, err = u.accountService.Get(tx, accountID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get account")
		return account.Account{}, make([]account_role.AccountRole, 0), err
	}

	roles, err = u.accountRoleService.GetAllByAccount(tx, accountID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get account's roles")
		return account.Account{}, make([]account_role.AccountRole, 0), err
	}

	acc.HashedPassword = ""
	return acc, roles, nil
}
