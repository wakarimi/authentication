package use_case

import (
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/model/account"
	"wakarimi-authentication/internal/model/account_role"
	"wakarimi-authentication/internal/model/device"
	"wakarimi-authentication/internal/model/refresh_token"
)

type transactor interface {
	WithTransaction(do func(tx *sqlx.Tx) (err error)) (err error)
}

type accessTokenService interface {
	Generate(refreshTokenID int, accountID int, deviceID int, roles []account_role.AccountRole) (string, error)
}

type accountRoleService interface {
	Assign(tx *sqlx.Tx, accountID int, accountRole account_role.RoleName) error
	GetAllByAccount(tx *sqlx.Tx, accountID int) ([]account_role.AccountRole, error)
	StringToRole(name string) (account_role.RoleName, error)
	HasRole(tx *sqlx.Tx, accountID int, roleName account_role.RoleName) (bool, error)
	Revoke(tx *sqlx.Tx, accountID int, accountRole account_role.RoleName) error
}

type accountService interface {
	IsUsernameTaken(tx *sqlx.Tx, username string) (bool, error)
	HashPassword(password string) (string, error)
	Create(tx *sqlx.Tx, account account.Account) (int, error)
	CountAccounts(tx *sqlx.Tx) (int, error)
	GetByUsername(tx *sqlx.Tx, username string) (account.Account, error)
	IsPasswordMatch(password string, hashedPassword string) bool
	UpdateLastSignIn(tx *sqlx.Tx, accountID int) error
	Get(tx *sqlx.Tx, accountID int) (account.Account, error)
	UpdatePassword(tx *sqlx.Tx, accountID int, hashedNewPassword string) error
	IsExists(tx *sqlx.Tx, accountID int) (bool, error)
	GetByIds(tx *sqlx.Tx, ids []int) ([]account.Account, error)
}

type deviceService interface {
	IsExistsByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (bool, error)
	GetByAccountAndFingerprint(tx *sqlx.Tx, accountID int, fingerprint string) (device.Device, error)
	Delete(tx *sqlx.Tx, deviceID int) error
	Create(tx *sqlx.Tx, create device.Device) (int, error)
	IsExists(tx *sqlx.Tx, deviceID int) (bool, error)
}

type refreshTokenService interface {
	DeleteByDevice(tx *sqlx.Tx, deviceID int) error
	GenerateAndCreateInDatabase(tx *sqlx.Tx, accountID int, deviceID int) (int, error)
	IsValid(tx *sqlx.Tx, token string) bool
	GetByToken(tx *sqlx.Tx, token string) (refresh_token.RefreshToken, error)
	GetPayload(token string) (refresh_token.Payload, error)
	Delete(tx *sqlx.Tx, refreshTokenID int) error
	DeleteByAccount(tx *sqlx.Tx, accountID int) error
	Get(tx *sqlx.Tx, refreshTokenID int) (refresh_token.RefreshToken, error)
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
