package account_service

import (
	"authentication/internal/errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (s Service) UpdateLastLogIn(tx *sqlx.Tx, accountID int) (err error) {
	isExists, err := s.IsExists(tx, accountID)
	if err != nil {
		return err
	}
	if !isExists {
		err = errors.NotFound{Resource: fmt.Sprintf("account with id=%d", accountID)}
		return err
	}

	err = s.AccountRepo.UpdateLastLogIn(tx, accountID)
	if err != nil {
		return err
	}

	return nil
}
