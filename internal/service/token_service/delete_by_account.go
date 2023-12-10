package token_service

import (
	"authentication/internal/errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) DeleteByAccount(tx *sqlx.Tx, accountID int) (err error) {
	log.Debug().Int("accountId", accountID).Msg("Deleting refresh token for account")

	isAccountExists, err := s.AccountService.IsExists(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to check account existence")
		return err
	}
	if !isAccountExists {
		err := errors.NotFound{Resource: fmt.Sprintf("account with id=%d", accountID)}
		log.Error().Err(err).Int("accountId", accountID).Msg("Account not found")
		return err
	}

	err = s.RefreshTokenRepo.DeleteByAccount(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to delete token for account")
		return err
	}

	log.Debug().Int("accountId", accountID).Msg("Token by account deleted")
	return nil
}
