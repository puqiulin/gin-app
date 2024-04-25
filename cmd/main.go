package main

import (
	"gin-app/pkg/logger"
	"gin-app/pkg/wire"
)

func main() {
	app, cleanup, err := wire.InitApp()
	defer cleanup()

	if err != nil {
		logger.Log.Fatalf("Init app error: %v", err)
		return
	}

	if err = app.Run(); err != nil {
		logger.Log.Fatalf("Run app error: %v", err)
		panic(err)
	}
}
