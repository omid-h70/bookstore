package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/omid-h70/bookstore/users-api/domain"
	"github.com/omid-h70/bookstore/users-api/services"
	"net/http"
)

//TODO in Controller Do The Parsing and Pass it To Service
//a Controller Doesnt Know how a User is Created , that's the point of why Services Exist

func CreateUser(c echo.Context) error {
	/* checking not needed maybe
	bytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}*/
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	newUser, err := services.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	return c.JSON(http.StatusCreated, newUser)
}

func GetUser(c echo.Context) error {
	//
	//	u := domain.User{}
	//	services.CreateUser(u)
	//
	return c.String(http.StatusNotImplemented, "implement me !")
}

func SearchUser(c echo.Context) error {
	//
	//	u := domain.User{}
	//	services.CreateUser(u)
	//
	return c.String(http.StatusNotImplemented, "implement me !")
}
