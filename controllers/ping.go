package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handlers) Pong(c echo.Context) error {
	return c.String(http.StatusOK, "Pong")
}
