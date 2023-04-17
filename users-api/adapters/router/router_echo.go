package router

import (
	"github.com/labstack/echo/v4"
	"github.com/omid-h70/bookstore/users-api/controllers"
)

type EchoRouter struct {
	echoMux *echo.Echo
}

func (e *EchoRouter) NewRouter() (Router, error) {
	router := &EchoRouter{
		echoMux: echo.New(),
	}
	return router, nil
}

func (e *EchoRouter) SetAppHandlers() {
	e.echoMux.GET("/ping", controllers.Ping)
}
