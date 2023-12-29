package use_case

import "github.com/jmoiron/sqlx"

type transactor interface {
	WithTransaction(do func(tx *sqlx.Tx) (err error)) (err error)
}

type accessTokenService interface {
}

type accountRoleService interface {
}

type accountService interface {
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
	accountRoleService accountRoleService,
	accountService accountService,
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
