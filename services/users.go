package services

import (
	"github.com/omid-h70/bookstore/users-api/domain"
	"github.com/omid-h70/bookstore/users-api/models/db"
)

//A Service is Going to Act with Controller
//Gets the input from Action function and map them to Domain Objects
//And Do the Extra Work

func (service *Service) CreateUser(param db.CreateUserParam) (domain.User, error) {
	if err := domain.Validate(param); err != nil {
		return domain.User{}, err
	}
	return service.CreateUser(param)
}
