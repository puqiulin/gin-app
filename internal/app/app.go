package app

import (
	"fmt"
	"net/http"

	"gin-app/internal/handler"
	"gin-app/internal/middleware"
	"gin-app/internal/route"
	"gin-app/pkg/config"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var ProviderSet = wire.NewSet(NewApp)

type App struct {
	Instance *gin.Engine
	Config   *config.AppConfig
	logger   *logrus.Logger
}

func NewApp(c *config.AppConfig, l *logrus.Logger, uh *handler.UserHandler) *App {
	g := gin.New()

	g.Use(gin.Recovery())
	g.Use(middleware.Logger())

	g.Use(static.Serve("/", static.LocalFile("./web/out", true)))

	//panic: '/api/user' in new path '/api/user' conflicts with existing wildcard '/*filepath' in existing prefix '/*filepath'
	//r.Static("/", "./web/.next")

	route.SetupRouter(g, uh)

	return &App{
		Instance: g,
		Config:   c,
		logger:   l,
	}
}

func (app *App) Run() error {
	port := app.Config.Port
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.Instance); err != nil {
		app.logger.Fatalf("Failed to start server: %v", err)
		return err
	}

	return nil
}
