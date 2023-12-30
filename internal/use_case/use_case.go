package use_case

import (
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/model/access_token"
	"wakarimi-authentication/internal/model/account"
	"wakarimi-authentication/internal/model/account_role"
	"wakarimi-authentication/internal/model/device"
	"wakarimi-authentication/internal/model/refresh_token"
)

type transactor interface {
	WithTransaction(do func(tx *sqlx.Tx) (err error)) (err error)
}

type accessTokenService interface {
	Generate(payload access_token.Payload) (string, error)
}

type accountRoleService interface {
	Assign(tx *sqlx.Tx, accountRole account_role.AccountRole) error
	GetAllByAccount(tx *sqlx.Tx, accountID int) ([]account_role.AccountRole, error)
}

type accountService interface {
	IsUsernameTaken(tx *sqlx.Tx, username string) (bool, error)
	HashPassword(password string) (string, error)
	Create(tx *sqlx.Tx, account account.Account) (int, error)
	CountAccounts(tx *sqlx.Tx) (int, error)
	GetByUsername(tx *sqlx.Tx, username string) (account.Account, error)
	IsPasswordMatch(password string, hashedPassword string) bool
	UpdateLastSignIn(tx *sqlx.Tx, accountID int) error
}

type deviceService interface {
	IsExistsByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (bool, error)
	GetByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (device.Device, error)
	Delete(tx *sqlx.Tx, deviceID int) error
	Create(tx *sqlx.Tx, create device.Device) (int, error)
}

type refreshTokenService interface {
	DeleteByDevice(tx *sqlx.Tx, deviceID int) error
	Generate(payload refresh_token.Payload) (string, error)
	Create(tx *sqlx.Tx, refreshToken refresh_token.RefreshToken) error
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
