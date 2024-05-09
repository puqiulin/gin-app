//go:build wireinject

package wire

import (
	"gin-app/internal/app"
	"gin-app/internal/handler"
	"gin-app/internal/repository"
	"gin-app/pkg/config"
	"gin-app/pkg/db"
	"gin-app/pkg/etcd"
	"gin-app/pkg/kafka"
	"gin-app/pkg/log"
	"github.com/google/wire"
)

func InitApp() (*app.App, func(), error) {
	panic(wire.Build(wire.NewSet(
		etcd.ProviderSet,
		config.ProviderSet,
		db.ProviderSet,
		repository.ProviderSet,
		handler.ProviderSet,
		log.ProviderSet,
		kafka.ProviderSet,
		app.ProviderSet,
	)))
}
