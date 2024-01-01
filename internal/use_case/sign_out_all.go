package use_case

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (u UseCase) SignOutAll(accountID int) (err error) {
	log.Debug().Int("accountId", accountID).Msg("Sign out by account")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = u.signOutAll(tx, accountID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to sign out")
		return err
	}

	return nil
}

func (u UseCase) signOutAll(tx *sqlx.Tx, accountID int) (err error) {
	err = u.refreshTokenService.DeleteByAccount(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to delete account's refresh tokens")
		return err
	}

	return nil
}
