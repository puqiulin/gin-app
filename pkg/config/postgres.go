package config

import (
	"gin-app/pkg/etcd"
	"gin-app/pkg/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string `json:"ssl_model"`
}

func NewPostgresConfig(client *clientv3.Client) (c *PostgresConfig, err error) {
	path := "/app/config/postgres"
	if c, err = etcd.LoadConfigFromPath[PostgresConfig](client, path); err != nil {
		logger.Log.Error("Loading etcd config error: %s", err)
	}
	return
}
