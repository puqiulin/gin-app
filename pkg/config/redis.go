package config

import (
	"gin-app/pkg/etcd"
	"gin-app/pkg/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type RedisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
}

func NewRedisConfig(client *clientv3.Client) (c *RedisConfig, err error) {
	path := "/app/config/redis"
	if c, err = etcd.LoadConfigFromPath[RedisConfig](client, path); err != nil {
		logger.Log.Error("Loading etcd config error: %s", err)
	}
	return
}
