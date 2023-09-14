package db

import "github.com/omid-h70/bookstore/users-api/domain"

type Store interface {
	GetUser(userId int64) (domain.User, error)
	CreateUser(CreateUserParam) (domain.User, error)
}
