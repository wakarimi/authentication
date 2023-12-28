package usecase

import (
	"context"
	"wakarimi-authentication/internal/service/token"
)

func (u *UseCase) VerifyToken(ctx context.Context) (isValid bool, accessTokenPayload token.AccessTokenPayload, err error) {
	return false, token.AccessTokenPayload{}, nil
}
