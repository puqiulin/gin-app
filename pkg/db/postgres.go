package db

import (
	"database/sql"
	"fmt"

	"gin-app/pkg/config"
	"gin-app/pkg/logger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"
)

func NewPostgres(c *config.PostgresConfig) (db *bun.DB, cleanup func(), err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
	//logger.Log.Infof("Postgres dsn is: %s", dsn)
	connector := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	db = bun.NewDB(sql.OpenDB(connector), pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName(connector.Config().Database),
	))

	if err = db.Ping(); err != nil {
		panic(err)
	}

	cleanup = func() {
		if err = db.Close(); err != nil {
			logger.Log.Error(err)
		}
	}

	return
}
