package redis

import (
	"context"
	"fmt"

	"gin-app/pkg/etcd"
	"gin-app/pkg/log"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewConfig, NewRedis)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
}

func NewConfig(client *clientv3.Client) (c *Config, err error) {
	path := "app/gin-app/config/redis"
	return etcd.LoadConfigFromPath[Config](client, path)
}

func NewRedis(c *Config) (db *redis.Client, cleanup func(), err error) {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Username: c.Username,
		Password: c.Password,
		DB:       c.DB,
	}
	db = redis.NewClient(opt)
	db.AddHook(redisotel.NewTracingHook())
	ctx, cancel := context.WithTimeout(context.Background(), opt.DialTimeout)
	defer cancel()

	if err = db.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	cleanup = func() {
		if err = db.Close(); err != nil {
			log.Log.Error(err)
		}
	}

	return
}
