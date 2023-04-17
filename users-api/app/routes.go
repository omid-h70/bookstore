package app

import (
	"github.com/labstack/echo/v4"
	"github.com/omid-h70/bookstore/users-api/controllers"
)

func makeRoutes(e *echo.Echo) {
	e.GET("/ping", controllers.Ping)
}
