package use_case

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/errors"
	"wakarimi-authentication/internal/model/account"
	"wakarimi-authentication/internal/model/account_role"
)

func (u UseCase) SignUp(username string, password string) error {
	log.Debug().Str("username", username).Msg("Creating an account")

	err := u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		isUsernameAlreadyTaken, err := u.accountService.IsUsernameTaken(tx, username)
		if err != nil {
			log.Error().Err(err).Str("username", username).Msg("Failed to check if the username is taken")
			return err
		}
		if isUsernameAlreadyTaken {
			err = errors.Conflict{Message: fmt.Sprintf("username %s is already taken", username)}
			log.Error().Err(err).Msg("Username is already taken")
			return err
		}

		hashedPassword, err := u.accountService.HashPassword(password)
		if err != nil {
			log.Error().Err(err).Msg("Failed to hash password")
			return err
		}

		accountToCreate := account.Account{
			Username:       username,
			HashedPassword: hashedPassword,
		}
		createdAccountID, err := u.accountService.Create(tx, accountToCreate)
		if err != nil {
			log.Error().Err(err).Str("username", username).Msg("Failed to create account")
			return err
		}

		numberOfAccounts, err := u.accountService.CountAccounts(tx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get the number of accounts")
			return err
		}
		if numberOfAccounts == 1 {
			accountRoleToAssign := account_role.AccountRole{
				AccountID: createdAccountID,
				Role:      account_role.RoleAdmin,
			}
			err := u.accountRoleService.Assign(tx, accountRoleToAssign)
			if err != nil {
				log.Error().Err(err).Msg("Failed to assign admin role to first user")
				return err
			}
		}

		accountRoleToAssign := account_role.AccountRole{
			AccountID: createdAccountID,
			Role:      account_role.RoleUser,
		}
		err = u.accountRoleService.Assign(tx, accountRoleToAssign)
		if err != nil {
			log.Error().Err(err).Msg("Failed to assign USER role")
			return err
		}

		log.Debug().Str("username", username).Int("accountId", createdAccountID).Msg("Account created")
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
