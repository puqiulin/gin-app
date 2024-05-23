package uptrace

import (
	"context"
	"time"

	"gin-app/pkg/config"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/uptrace-go/uptrace"
)

var ProviderSet = wire.NewSet(NewUptraceClient)

type UptraceClient struct{}

func NewUptraceClient(c *config.UptraceConfig, l *logrus.Logger) (*UptraceClient, func()) {
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN(c.DSN),
		uptrace.WithServiceName(c.ServiceName),
		uptrace.WithServiceVersion(c.ServiceVersion),
		uptrace.WithDeploymentEnvironment(c.Environment),
	)
	return &UptraceClient{}, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
		defer cancel()
		if err := uptrace.Shutdown(ctx); err != nil {
			l.Error("action", "uptrace.Shutdown", "err", err)
		}
	}
}
