package account_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (s Service) GetByUsername(tx *sqlx.Tx, username string) (account.Account, error) {
	log.Debug().Str("username", username).Msg("Getting account by username")

	readAccount, err := s.accountRepo.ReadByUsername(tx, username)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Failed to read account by username")
		return account.Account{}, err
	}

	log.Debug().Str("username", username).Int("accountId", readAccount.ID).Msg("Account got successfully")
	return readAccount, nil
}
