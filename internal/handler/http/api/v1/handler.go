package v1

import (
	"context"
	"github.com/rs/zerolog"
	"wakarimi-authentication/internal/model/account"
	"wakarimi-authentication/internal/service/token"
)

type UseCase interface {
	SignIn(ctx context.Context) (err error)
	SignOut(ctx context.Context) (err error)
	SignOutAll(ctx context.Context) (err error)

	RefreshToken(ctx context.Context) (newRefreshToken string, newAccessToken string, err error)
	VerifyToken(ctx context.Context) (isValid bool, accessTokenPayload token.AccessTokenPayload, err error)

	GetAccount(ctx context.Context) (account account.Account, err error)
	GetAccounts(ctx context.Context) (accounts []account.Account, err error)
	ChangePassword(ctx context.Context) (err error)
	SignUp(ctx context.Context) (err error)

	AssignRole(ctx context.Context) (err error)
	RevokeRole(ctx context.Context) (err error)
}

type Handler struct {
	uc     UseCase
	logger zerolog.Logger
}

func NewHandler(uc UseCase, logger zerolog.Logger) *Handler {
	return &Handler{
		uc:     uc,
		logger: logger,
	}
}
