package db

import (
	"context"
	"fmt"

	"gin-app/pkg/config"
	"gin-app/pkg/logger"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

func NewRedis(c *config.RedisConfig) (db *redis.Client, cleanup func(), err error) {
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
			logger.Log.Error(err)
		}
	}

	return
}
