package account

import "context"

type Repository interface {
	Create(ctx context.Context, account Account) (accountID int, err error)
	Read(ctx context.Context, accountID int) (account Account, err error)
	ReadByUsername(ctx context.Context, username string) (account Account, err error)
	UpdateLastLogin(ctx context.Context, accountID int) (err error)
	UpdatePassword(ctx context.Context, accountID int, hashedPassword string) (err error)
	IsExists(ctx context.Context, accountID int) (exists bool, err error)
	IsUsernameTaken(ctx context.Context, username string) (taken bool, err error)
	CountAccounts(ctx context.Context) (count int, err error)
}
