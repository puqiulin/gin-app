package kafka

import (
	"fmt"

	"gin-app/pkg/etcd"
	"github.com/google/wire"
	"github.com/segmentio/kafka-go"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewConfig, NewKafkaWriter)

type Config struct {
	Host  string
	Port  int
	Topic string
}

func NewConfig(client *clientv3.Client) (c *Config, err error) {
	path := "app/gin-app/config/kafka"
	return etcd.LoadConfigFromPath[Config](client, path)
}

func NewKafkaWriter(c *Config) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(fmt.Sprintf("%s:%d", c.Host, c.Port)),
		Topic:                  c.Topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
}
