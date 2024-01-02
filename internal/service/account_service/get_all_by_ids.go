package account_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (s Service) GetByIds(tx *sqlx.Tx, ids []int) (readAccounts []account.Account, err error) {
	log.Debug().Msg("Getting account by username")

	readAccounts, err = s.accountRepo.ReadAllByIds(tx, ids)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read account by username")
		return make([]account.Account, 0), err
	}

	log.Debug().Int("count", len(readAccounts)).Msg("Account got successfully")
	return readAccounts, nil
}
