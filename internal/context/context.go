package context

import (
	"authentication/internal/config"
	"github.com/jmoiron/sqlx"
)

type AppContext struct {
	Config *config.Configuration
	Db     *sqlx.DB
}
