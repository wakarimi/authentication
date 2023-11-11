package account_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) IsUsernameTaken(tx *sqlx.Tx, username string) (isAlreadyTaken bool, err error) {
	log.Debug().Str("username", username).Msg("Checking whether the username is already being used")

	isAlreadyTaken, err = s.AccountRepo.IsUsernameTaken(tx, username)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Failed to check if the username is taken")
		return false, err
	}

	log.Debug().Str("username", username).Bool("isAlreadyTaken", isAlreadyTaken).Msg("The usage of the user name has been checked")
	return isAlreadyTaken, nil
}
