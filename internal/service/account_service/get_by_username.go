package account_service

import (
	"authentication/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) GetByUsername(tx *sqlx.Tx, username string) (account model.Account, err error) {
	account, err = s.AccountRepo.ReadByUsername(tx, username)
	if err != nil {
		log.Error().Err(err).Str("username", username).Msg("Failed to read account by username")
		return model.Account{}, err
	}

	return account, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
