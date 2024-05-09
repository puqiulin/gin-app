package config

import (
	"gin-app/pkg/etcd"
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
	path := "app/gin-app/config/redis"
	return etcd.LoadConfigFromPath[RedisConfig](client, path)
}
