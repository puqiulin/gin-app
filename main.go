package main

import (
	"gin-app/pkg/wire"
)

func main() {
	app, cleanup, err := wire.InitApp()

	defer cleanup()
	if err != nil {
		panic(err)
	}

	if err = app.Run(); err != nil {
		panic(err)
	}
}
