package kafka

import (
	"fmt"

	"gin-app/pkg/config"
	"github.com/google/wire"
	"github.com/segmentio/kafka-go"
)

var ProviderSet = wire.NewSet(NewKafkaWriter)

func NewKafkaWriter(c *config.KafkaConfig) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(fmt.Sprintf("%s:%d", c.Host, c.Port)),
		Topic:                  c.Topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
}
