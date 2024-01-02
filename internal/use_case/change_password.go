package use_case

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/errors"
)

func (u UseCase) ChangePassword(accountID int, oldPassword string, newPassword string) (err error) {
	log.Debug().Msg("Changing password")

	err = u.transactor.WithTransaction(func(tx *sqlx.Tx) (err error) {
		err = u.changePassword(tx, accountID, oldPassword, newPassword)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to change password")
		return err
	}

	return nil
}

func (u UseCase) changePassword(tx *sqlx.Tx, accountID int, oldPassword string, newPassword string) (err error) {
	isAccountExists, err := u.accountService.IsExists(tx, accountID)
	if err != nil {
		log.Error().Err(err).Int("accountId", accountID).Msg("Failed to check account existence")
		return err
	}
	if !isAccountExists {
		err = errors.NotFound{EntityName: fmt.Sprintf("account with id=%d", accountID)}
		log.Error().Err(err).Int("accountId", accountID).Msg("Account not found")
		return err
	}
	account, err := u.accountService.Get(tx, accountID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read account by username")
		return err
	}
	isPasswordMatch := u.accountService.IsPasswordMatch(oldPassword, account.HashedPassword)
	if !isPasswordMatch {
		err := errors.Unauthorized{Message: "invalid old password"}
		log.Error().Err(err).Int("accountId", account.ID).Str("username", account.Username).Msg("Invalid old password")
		return err
	}

	hashedNewPassword, err := u.accountService.HashPassword(newPassword)
	if err != nil {
		log.Error().Err(err).Msg("Failed to hash password")
		return err
	}

	err = u.accountService.UpdatePassword(tx, account.ID, hashedNewPassword)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update password")
		return err
	}

	err = u.refreshTokenService.DeleteByAccount(tx, account.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete account's tokens")
		return err
	}

	log.Debug().Msg("Password changed")
	return nil
}
