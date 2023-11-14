package account_service

import (
	"authentication/internal/errors"
	"authentication/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) Create(tx *sqlx.Tx, username string, password string) (err error) {
	log.Debug().Str("username", username).Msg("Creating an account")

	isUsernameAlreadyTaken, err := s.IsUsernameTaken(tx, username)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Failed to check if the username is taken")
		return err
	}
	if isUsernameAlreadyTaken {
		err = errors.Conflict{Message: fmt.Sprintf("username %s is already taken", username)}
		log.Error().Err(err).Msg("Username is already taken")
		return err
	}

	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		log.Error().Msg("Failed to hash password")
		return err
	}

	accountToCreate := model.Account{
		Username:       username,
		HashedPassword: hashedPassword,
	}
	createdAccountID, err := s.AccountRepo.Create(tx, accountToCreate)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Failed to create an account")
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

	log.Debug().Int("accountID", createdAccountID).Str("username", username).Msg("Account created")
	return nil
}
