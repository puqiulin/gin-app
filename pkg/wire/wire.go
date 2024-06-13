//go:build wireinject

package wire

import (
	"gin-app/internal/graphql"
	"gin-app/internal/handler"
	"gin-app/internal/repository"
	"gin-app/pkg/app"
	"gin-app/pkg/db/postgres"
	"gin-app/pkg/db/redis"
	"gin-app/pkg/etcd"
	"gin-app/pkg/kafka"
	"gin-app/pkg/log"
	"gin-app/pkg/uptrace"
	"github.com/google/wire"
)

func InitApp() (*app.App, func(), error) {
	panic(wire.Build(wire.NewSet(
		app.ProviderSet,
		etcd.ProviderSet,
		postgres.ProviderSet,
		redis.ProviderSet,
		uptrace.ProviderSet,
		repository.ProviderSet,
		handler.ProviderSet,
		graphql.ProviderSet,
		log.ProviderSet,
		kafka.ProviderSet,
	)))
}
