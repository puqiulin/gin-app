package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

var ProviderSet = wire.NewSet(NewRepository)

type Repository struct {
	db  *bun.DB
	rdb *redis.Client
	l   *logrus.Logger
}

func NewRepository(db *bun.DB, rdb *redis.Client, l *logrus.Logger) *Repository {
	l.WithField("repository", "user")
	return &Repository{db: db, rdb: rdb, l: l}
}
