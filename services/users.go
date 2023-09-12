package services

import "github.com/omid-h70/bookstore/users-api/domain"

//A Service is Going to Act with Controller
//Gets the input from Action function and map them to Domain Objects
//And Do the Extra Work

func CreateUser(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
