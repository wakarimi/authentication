package use_case

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (u UseCase) GetAccounts(ids []int) (accs []account.Account, err error) {
	log.Debug().Msg("Getting accounts")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		accs, err = u.getAccounts(tx, ids)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get account details")
		return make([]account.Account, 0), err
	}

	return accs, nil
}

func (u UseCase) getAccounts(tx *sqlx.Tx, ids []int) (accs []account.Account, err error) {
	accs, err = u.accountService.GetByIds(tx, ids)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get accounts")
		return make([]account.Account, 0), err
	}
	return accs, nil
}
