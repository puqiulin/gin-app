package log

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

// KafkaHook implements logrus Hook interface.
type KafkaHook struct {
	Writer *kafka.Writer
}

// Levels returns the levels that this hook is interested in.
func (hook *KafkaHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire is called by Logrus when something is logged.
func (hook *KafkaHook) Fire(entry *logrus.Entry) error {
	msg, err := entry.String()
	if err != nil {
		return err
	}
	return hook.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(entry.Level.String()),
			Value: []byte(msg),
		},
	)
}
