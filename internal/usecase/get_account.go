package usecase

import (
	"context"
	"wakarimi-authentication/internal/model/account"
)

func (u *UseCase) GetAccount(ctx context.Context) (acc account.Account, err error) {
	return account.Account{}, nil
}
