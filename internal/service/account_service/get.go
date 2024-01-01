package account_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (s Service) Get(tx *sqlx.Tx, accountID int) (account.Account, error) {
	log.Debug().Msg("Getting account by username")

	readAccount, err := s.accountRepo.Read(tx, accountID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read account by username")
		return account.Account{}, err
	}

	log.Debug().Int("accountId", readAccount.ID).Msg("Account got successfully")
	return readAccount, nil
}
