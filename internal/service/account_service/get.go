package account_service

import (
	"authentication/internal/errors"
	"authentication/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (s Service) Get(tx *sqlx.Tx, accountID int) (account model.Account, err error) {
	isExists, err := s.IsExists(tx, accountID)
	if err != nil {
		return model.Account{}, err
	}
	if !isExists {
		err = errors.NotFound{Resource: fmt.Sprintf("account with id=%d", accountID)}
		return model.Account{}, err
	}

	account, err = s.AccountRepo.Read(tx, accountID)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}
