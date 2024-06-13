package uptrace

import (
	"context"
	"time"

	"gin-app/pkg/etcd"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/uptrace-go/uptrace"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewClient, NewConfig)

type Client struct {
	//Tracer *trace.Tracer
}

type Config struct {
	DSN            string
	ServiceName    string
	ServiceVersion string
	Environment    string
}

func NewConfig(client *clientv3.Client) (c *Config, err error) {
	path := "app/gin-app/config/uptrace"
	return etcd.LoadConfigFromPath[Config](client, path)
}

// NewClient https://uptrace.dev/get/opentelemetry-go.html#quickstart
func NewClient(c *Config, l *logrus.Logger) (*Client, func()) {
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN(c.DSN),
		uptrace.WithServiceName(c.ServiceName),
		uptrace.WithServiceVersion(c.ServiceVersion),
		uptrace.WithDeploymentEnvironment(c.Environment),
	)

	return &Client{
			// Create a tracer. Usually, tracer is a global variable.
			//Tracer: otel.Tracer(c.ServiceName),
		},
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
			defer cancel()
			if err := uptrace.Shutdown(ctx); err != nil {
				l.Error("uptrace.Shutdown", "err", err)
			}
		}
}
