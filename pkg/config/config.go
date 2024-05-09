package config

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewAppConfig, NewRedisConfig, NewPostgresConfig, NewKafkaConfig)
