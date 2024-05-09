package config

import (
	"gin-app/pkg/etcd"
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
	path := "app/gin-app/config/postgres"
	return etcd.LoadConfigFromPath[PostgresConfig](client, path)
}
