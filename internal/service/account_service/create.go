package account_service

import (
	"authentication/internal/errors"
	"authentication/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) Create(tx *sqlx.Tx, account model.Account, password string) (err error) {
	log.Debug().Str("usernme", account.Username).Msg("Account creation")

	isUsernameTaken, err := s.IsUsernameTaken(tx, account.Username)
	if err != nil {
		log.Error().Err(err).Str("username", account.Username).Msg("Failed to check if the username is taken")
		return err
	}
	if isUsernameTaken {
		err = errors.Conflict{Message: fmt.Sprintf("username %s is already taken", account.Username)}
		log.Error().Err(err).Msg("Username is already taken")
		return err
	}

	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		log.Error().Msg("Failed to hash password")
		return err
	}

	account.HashedPassword = hashedPassword

	createdAccountID, err := s.AccountRepo.Create(tx, account)
	if err != nil {
		log.Error().Err(err).Str("username", account.Username).Msg("Failed to create account")
		return err
	}

	numberOfAccounts, err := s.CountAccounts(tx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get the number of accounts")
	}
	if numberOfAccounts == 1 {
		err = s.AccountRoleService.AssignRole(tx, createdAccountID, model.RoleAdmin)
		if err != nil {
			log.Error().Err(err).Int("accountId", createdAccountID).Msg("Failed to assign ADMIN role to account")
			return err
		}
	}
	err = s.AccountRoleService.AssignRole(tx, createdAccountID, model.RoleUser)
	if err != nil {
		log.Error().Err(err).Int("accountId", createdAccountID).Msg("Failed to assign USER role to account")
		return err
	}

	log.Debug().Int("accountId", createdAccountID).Str("username", account.Username).Msg("Account created")
	return nil
}
