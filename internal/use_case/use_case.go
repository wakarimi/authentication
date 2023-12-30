package use_case

import (
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/model/account"
	"wakarimi-authentication/internal/model/account_role"
)

type transactor interface {
	WithTransaction(do func(tx *sqlx.Tx) (err error)) (err error)
}

type accessTokenService interface {
}

type accountRoleService interface {
	Assign(tx *sqlx.Tx, accountRole account_role.AccountRole) error
}

type accountService interface {
	IsUsernameTaken(tx *sqlx.Tx, username string) (bool, error)
	HashPassword(password string) (string, error)
	Create(tx *sqlx.Tx, account account.Account) (int, error)
	CountAccounts(tx *sqlx.Tx) (int, error)
}

type deviceService interface {
}

type refreshTokenService interface {
}

type UseCase struct {
	transactor          transactor
	accessTokenService  accessTokenService
	accountRoleService  accountRoleService
	accountService      accountService
	deviceService       deviceService
	refreshTokenService refreshTokenService
}

func New(transactor transactor,
	accessTokenService accessTokenService,
	accountService accountService,
	accountRoleService accountRoleService,
	deviceService deviceService,
	refreshTokenService refreshTokenService) *UseCase {
	return &UseCase{
		transactor:          transactor,
		accountRoleService:  accountRoleService,
		accountService:      accountService,
		deviceService:       deviceService,
		refreshTokenService: refreshTokenService,
		accessTokenService:  accessTokenService,
	}
}
