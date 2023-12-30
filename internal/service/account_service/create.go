package account_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (s Service) Create(tx *sqlx.Tx, account account.Account) (int, error) {
	log.Debug().Str("username", account.Username).Msg("Creating an account")

	accountID, err := s.accountRepo.Create(tx, account)
	if err != nil {
		log.Error().Err(err).Str("username", account.Username).Msg("Failed to create account")
		return 0, err
	}

	log.Debug().Str("username", account.Username).Int("accountId", accountID).Msg("Account created")
	return accountID, nil
}
