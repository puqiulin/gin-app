// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
)

// Injectors from wire.go:

func InitApp() (*app.App, func(), error) {
	client := etcd.NewClient()
	config, err := app.NewConfig(client)
	if err != nil {
		return nil, nil, err
	}
	uptraceConfig, err := uptrace.NewConfig(client)
	if err != nil {
		return nil, nil, err
	}
	kafkaConfig, err := kafka.NewConfig(client)
	if err != nil {
		return nil, nil, err
	}
	writer := kafka.NewKafkaWriter(kafkaConfig)
	logger := log.NewLog(writer)
	uptraceClient, cleanup := uptrace.NewClient(uptraceConfig, logger)
	postgresConfig, err := postgres.NewConfig(client)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	db, cleanup2, err := postgres.NewPostgres(postgresConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	redisConfig, err := redis.NewConfig(client)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	redisClient, cleanup3, err := redis.NewRedis(redisConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	repositoryRepository := repository.NewRepository(db, redisClient, logger)
	userHandler := handler.NewUserHandler(repositoryRepository, logger)
	userResolver := graphql.NewUserResolver(repositoryRepository, logger)
	userField := graphql.NewUserField(userResolver)
	postResolver := graphql.NewPostResolver(repositoryRepository, logger)
	postField := graphql.NewPostField(postResolver)
	rootSchema := graphql.NewRootSchema(userField, postField)
	graphQLHandler := handler.NewGraphQLHandler(rootSchema, logger)
	googleHandler := handler.NewGoogleHandler(repositoryRepository, logger)
	appApp := app.NewApp(config, uptraceConfig, uptraceClient, logger, userHandler, graphQLHandler, googleHandler)
	return appApp, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
