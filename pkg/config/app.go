package config

import (
	"gin-app/pkg/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type AppConfig struct {
	Port int
}

func NewAppConfig(client *clientv3.Client) (c *AppConfig, err error) {
	path := "app/gin-app/config/app"
	return etcd.LoadConfigFromPath[AppConfig](client, path)
}
