package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"wakarimi-authentication/internal/service/account_role_service"
	"wakarimi-authentication/internal/service/account_service"
	"wakarimi-authentication/internal/service/device_service"
	"wakarimi-authentication/internal/service/refresh_token_service"
	"wakarimi-authentication/internal/store/pgsqlx_repo/account_repo"
	"wakarimi-authentication/internal/store/pgsqlx_repo/account_role_repo"
	"wakarimi-authentication/internal/store/pgsqlx_repo/device_repo"
	"wakarimi-authentication/internal/store/pgsqlx_repo/refresh_token_repo"
	"wakarimi-authentication/internal/usecase"
)

type Container struct {
	pgSqlx *sqlx.DB
	http   *gin.Engine
}

func NewContainer(pgSqlxConnection *sqlx.DB) *Container {
	return &Container{
		pgSqlx: pgSqlxConnection,
	}
}

func (c *Container) GetUseCase() *usecase.UseCase {
	return usecase.NewUseCase(
		c.getAccountRoleService(),
		c.getAccountService(),
		c.getDeviceService(),
		c.getRefreshTokenService(),
	)
}

func (c *Container) getPgSqlx() *sqlx.DB {
	return c.pgSqlx
}

func (c *Container) getHttp() *gin.Engine {
	return c.http
}

func (c *Container) getAccountRoleService() *account_role_service.Service {
	return account_role_service.NewService(c.getAccountRoleRepo())
}

func (c *Container) getAccountService() *account_service.Service {
	return account_service.NewService(c.getAccountRepo())
}

func (c *Container) getDeviceService() *device_service.Service {
	return device_service.NewService(c.getDeviceRepo())
}

func (c *Container) getRefreshTokenService() *refresh_token_service.Service {
	return refresh_token_service.NewService(c.getRefreshTokenRepo())
}

func (c *Container) getAccountRoleRepo() *account_role_repo.Repository {
	return account_role_repo.NewPgSQLRepository(c.pgSqlx)
}

func (c *Container) getAccountRepo() *account_repo.Repository {
	return account_repo.NewPgSQLRepository(c.pgSqlx)
}

func (c *Container) getDeviceRepo() *device_repo.Repository {
	return device_repo.NewPgSQLRepository(c.pgSqlx)
}

func (c *Container) getRefreshTokenRepo() *refresh_token_repo.Repository {
	return refresh_token_repo.NewPgSQLRepository(c.pgSqlx)
}
