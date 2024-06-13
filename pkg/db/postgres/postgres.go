package postgres

import (
	"database/sql"
	"fmt"

	"gin-app/pkg/etcd"
	"gin-app/pkg/log"
	"github.com/google/wire"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewConfig, NewPostgres)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string `json:"ssl_model"`
}

func NewConfig(client *clientv3.Client) (c *Config, err error) {
	path := "app/gin-app/config/postgres"
	return etcd.LoadConfigFromPath[Config](client, path)
}

func NewPostgres(c *Config) (db *bun.DB, cleanup func(), err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
	//log.Log.Infof("Postgres dsn is: %s", dsn)
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
			log.Log.Error(err)
		}
	}

	return
}
