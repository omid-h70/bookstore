package app

import (
	"github.com/labstack/echo/v4"
	"github.com/omid-h70/bookstore/users-api/controllers"
)

func (app *App) setRoutes(e *echo.Echo) {

	h := controllers.NewHandler(app.store)
	//readiness
	e.GET("/ping", h.Pong)

	//Users
	e.GET("/search", h.SearchUser)
	e.GET("/users/:user_id", h.GetUser)
	e.POST("/users", h.CreateUser)

}
