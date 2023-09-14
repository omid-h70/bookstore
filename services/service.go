package services

import "github.com/omid-h70/bookstore/users-api/models/db"

type Service struct {
	store *db.Store
}

func NewService(store *db.Store) Service {
	return Service{
		store: store,
	}
}
