package app

import (
	"fmt"
	"net/http"
	"time"

	myHandler "gin-app/internal/handler"
	"gin-app/internal/middleware"
	"gin-app/internal/route"
	"gin-app/pkg/etcd"
	"gin-app/pkg/uptrace"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

var ProviderSet = wire.NewSet(NewConfig, NewApp)

type Config struct {
	Port int
}

func NewConfig(client *clientv3.Client) (c *Config, err error) {
	path := "app/gin-app/config/app"
	return etcd.LoadConfigFromPath[Config](client, path)
}

type App struct {
	Instance      *gin.Engine
	AppConfig     *Config
	UptraceConfig *uptrace.Config
	UptraceClient *uptrace.Client
	logger        *logrus.Logger
}

func NewApp(
	ca *Config,
	uco *uptrace.Config,
	ucl *uptrace.Client,
	l *logrus.Logger,
	uh *myHandler.UserHandler,
	gh *myHandler.GraphQLHandler,
	goh *myHandler.GoogleHandler,
) *App {
	g := gin.New()

	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://127.0.0.1:3001", "http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//g.Use(gin.Recovery())
	g.Use(middleware.Logger())

	//uptrace
	//https://uptrace.dev/get/instrument/opentelemetry-gin.html#gin-instrumentation
	g.Use(otelgin.Middleware(uco.ServiceName))

	g.Use(static.Serve("/", static.LocalFile("./web/out", true)))

	//panic: '/api/user' in new path '/api/user' conflicts with existing wildcard '/*filepath' in existing prefix '/*filepath'
	//r.Static("/", "./web/.next")

	route.SetupRouter(g, uh, gh, goh)

	return &App{
		Instance:      g,
		AppConfig:     ca,
		UptraceConfig: uco,
		UptraceClient: ucl,
		logger:        l,
	}
}

func (app *App) Run() error {
	port := app.AppConfig.Port
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.Instance); err != nil {
		app.logger.Fatalf("Failed to start server: %v", err)
		return err
	}

	return nil
}
