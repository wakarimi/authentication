package account

import (
	"context"
	"wakarimi-authentication/internal/service"
)

type Repository interface {
	service.Transactor
	Create(ctx context.Context, account Account) (accountID int, err error)
}
