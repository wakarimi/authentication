package usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"wakarimi-authentication/internal/model/account"
)

func (u *UseCase) SignUp(ctx context.Context) (a account.Account, err error) {
	log.Warn().Msg("Implement SignUp UseCase")
	return account.Account{}, nil
}
