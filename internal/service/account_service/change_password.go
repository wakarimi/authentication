package account_service

import (
	"authentication/internal/errors"
	"authentication/internal/utils"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (s Service) ChangePassword(tx *sqlx.Tx, accountID int, oldPassword string, newPassword string) (err error) {
	log.Debug().Int("accountId", accountID).Msg("Account creation")

	account, err := s.Get(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to get account by username")
		return err
	}
	isMatch := utils.CheckPasswordHash(oldPassword, account.HashedPassword)
	if !isMatch {
		err := errors.Unauthorized{Message: "invalid old password"}
		log.Error().Err(err).Int("accountId", accountID).Msg("Invalid old password")
		return err
	}

	hashedNewPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		log.Error().Msg("Failed to hash password")
		return err
	}

	err = s.AccountRepo.ChangePassword(tx, accountID, hashedNewPassword)
	if err != nil {
		log.Error().Err(err).Str("username", account.Username).Msg("Failed to change password")
		return err
	}

	log.Debug().Str("username", account.Username).Msg("Password changed")
	return nil
}
