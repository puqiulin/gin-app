package config

import (
	"gin-app/pkg/etcd"
	"gin-app/pkg/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type AppConfig struct {
	Port int
}

func NewAppConfig(client *clientv3.Client) (c *AppConfig, err error) {
	path := "/app/config/app"
	if c, err = etcd.LoadConfigFromPath[AppConfig](client, path); err != nil {
		logger.Log.Error("Loading etcd config error: %s", err)
	}
	return
}
