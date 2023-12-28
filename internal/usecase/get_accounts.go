package usecase

import (
	"context"
	"wakarimi-authentication/internal/model/account"
)

func (u *UseCase) GetAccounts(ctx context.Context) (accs []account.Account, err error) {
	return make([]account.Account, 0), nil
}
