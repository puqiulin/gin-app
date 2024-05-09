package config

import (
	"gin-app/pkg/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type KafkaConfig struct {
	Host  string
	Port  int
	Topic string
}

func NewKafkaConfig(client *clientv3.Client) (c *KafkaConfig, err error) {
	path := "app/gin-app/config/kafka"
	return etcd.LoadConfigFromPath[KafkaConfig](client, path)
}
