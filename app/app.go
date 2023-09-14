package app

import (
	echo "github.com/labstack/echo/v4"
	"github.com/omid-h70/bookstore/users-api/models/db"
	"github.com/omid-h70/bookstore/users-api/utils"
)

type App struct {
	mux    *echo.Echo
	config *utils.Config
	store  *db.Store
}

func NewApp(config *utils.Config, echoMux *echo.Echo) App {
	return App{
		config: config,
		mux:    echoMux,
	}
}

func (app *App) Run() error {
	return app.mux.Start(app.config.HttpServerAddress)
}
