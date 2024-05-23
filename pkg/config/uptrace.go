package config

import (
	"gin-app/pkg/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type UptraceConfig struct {
	DSN            string
	ServiceName    string
	ServiceVersion string
	Environment    string
}

func NewUptraceConfig(client *clientv3.Client) (c *UptraceConfig, err error) {
	path := "app/gin-app/config/uptrace"
	return etcd.LoadConfigFromPath[UptraceConfig](client, path)
}
