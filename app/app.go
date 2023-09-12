package app

import (
	echo "github.com/labstack/echo/v4"
	"github.com/omid-h70/bookstore/users-api/util"
)

type App struct {
	mux    *echo.Echo
	config *util.Config
}

func NewApp(config *util.Config, echoMux *echo.Echo) App {
	return App{
		config: config,
		mux:    echoMux,
	}
}

func (app *App) Run() error {
	return app.mux.Start(app.config.HttpServerAddress)
}
