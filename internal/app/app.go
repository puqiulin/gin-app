package app

import (
	"fmt"
	"net/http"

	"gin-app/internal/handler"
	"gin-app/internal/middleware"
	"gin-app/internal/route"
	"gin-app/pkg/config"
	"gin-app/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewApp)

type App struct {
	Instance *gin.Engine
	Config   *config.AppConfig
}

func NewApp(c *config.AppConfig, uh *handler.UserHandler) *App {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	r.Static("/static", "./static")

	route.SetupRouter(r, uh)

	return &App{
		Instance: r,
		Config:   c,
	}
}

func (app *App) Run() error {
	port := app.Config.Port
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.Instance); err != nil {
		logger.Log.Fatalf("Failed to start server: %v", err)
		return err
	}

	logger.Log.Infof("Serer running at port: %d...", port)
	return nil
}
