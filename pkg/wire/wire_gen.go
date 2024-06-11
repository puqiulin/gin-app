// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"gin-app/internal/app"
	"gin-app/internal/graphql"
	"gin-app/internal/handler"
	"gin-app/internal/repository"
	"gin-app/pkg/config"
	"gin-app/pkg/db"
	"gin-app/pkg/etcd"
	"gin-app/pkg/kafka"
	"gin-app/pkg/log"
)

// Injectors from wire.go:

func InitApp() (*app.App, func(), error) {
	client := etcd.NewEtcdClient()
	appConfig, err := config.NewAppConfig(client)
	if err != nil {
		return nil, nil, err
	}
	kafkaConfig, err := config.NewKafkaConfig(client)
	if err != nil {
		return nil, nil, err
	}
	writer := kafka.NewKafkaWriter(kafkaConfig)
	logger := log.NewLog(writer)
	postgresConfig, err := config.NewPostgresConfig(client)
	if err != nil {
		return nil, nil, err
	}
	bunDB, cleanup, err := db.NewPostgres(postgresConfig)
	if err != nil {
		return nil, nil, err
	}
	redisConfig, err := config.NewRedisConfig(client)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	redisClient, cleanup2, err := db.NewRedis(redisConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	repositoryRepository := repository.NewRepository(bunDB, redisClient, logger)
	userHandler := handler.NewUserHandler(repositoryRepository, logger)
	userResolver := graphql.NewUserResolver(repositoryRepository, logger)
	userField := graphql.NewUserField(userResolver)
	postResolver := graphql.NewPostResolver(repositoryRepository, logger)
	postField := graphql.NewPostField(postResolver)
	rootSchema := graphql.NewRootSchema(userField, postField)
	graphQLHandler := handler.NewGraphQLHandler(rootSchema, logger)
	googleHandler := handler.NewGoogleHandler(repositoryRepository, logger)
	appApp := app.NewApp(appConfig, logger, userHandler, graphQLHandler, googleHandler)
	return appApp, func() {
		cleanup2()
		cleanup()
	}, nil
}
