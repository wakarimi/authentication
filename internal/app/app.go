package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"wakarimi-authentication/internal/config"
	"wakarimi-authentication/internal/handler/http"
	"wakarimi-authentication/internal/store/pgsqlx_repo"
)

type App struct {
	logger     zerolog.Logger
	cfg        config.Config
	pgSqlx     *sqlx.DB
	httpServer *http.Server
	container  *Container
	router     *http.Router
}

func NewApp(configPath string) (app *App, err error) {
	app = &App{}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	app.cfg = cfg

	app.initLogger(app.cfg.App)

	pgSqlxConnection, err := pgsqlx_repo.InitDB(cfg.DB)
	if err != nil {
		return nil, err
	}
	app.pgSqlx = pgSqlxConnection

	app.router = http.NewRouter()

	httpServer := http.NewServer(cfg.HTTP, app.router.Engine())
	app.httpServer = httpServer

	app.container = NewContainer(app.pgSqlx)

	return app, nil
}
