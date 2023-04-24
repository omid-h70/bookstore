package controllers

import (
	"github.com/omid-h70/bookstore/users-api/adapters/controller"
	"github.com/omid-h70/bookstore/users-api/domain"
	"github.com/omid-h70/bookstore/users-api/services"
)

//TODO in Controller Do The Parsing and Pass it To Service
//a Controller Doesnt Know how a User is Created , that's the point of why Services Exist

func CreateUser(request *controller.Request) controller.Response {

	u := domain.User{}
	services.CreateUser(u)

	return controller.Response{}
}
