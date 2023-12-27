package v1

import (
	"context"
	"github.com/rs/zerolog"
	"wakarimi-authentication/internal/model/account"
)

type UseCase interface {
	SignUp(ctx context.Context) (account account.Account, err error)
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
