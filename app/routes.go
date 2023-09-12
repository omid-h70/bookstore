package app

import (
	"github.com/labstack/echo/v4"
	"github.com/omid-h70/bookstore/users-api/controllers"
)

func (app *App) setRoutes(e *echo.Echo) {
	//readiness
	e.GET("/ping", controllers.Pong)

	//Users
	e.GET("/search", controllers.SearchUser)
	e.GET("/users/:user_id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)

}
