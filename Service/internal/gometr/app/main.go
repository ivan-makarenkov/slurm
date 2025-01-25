package app

import (
	"log"

	"gotemplate/Service/internal/gometr/handlers"
	"gotemplate/Service/internal/gometr/infrastructure/config"
	"gotemplate/Service/pkg/graceful"

	"go.uber.org/zap"
)

type App struct {
	log     *zap.Logger
	cfg     *config.Config
	handler *handlers.Handler
}

func Start() {
	app := new(App)
	if err := app.Bootstrap(); err != nil {
		log.Fatal(err)
	}

	go func() {
		app.Run()
	}()

	err := graceful.WaitShutdown()
	if err != nil {
		app.log.Fatal("gometr is dead")
	} else {
		app.log.Info("gometr gracefully stopped")
	}
}
