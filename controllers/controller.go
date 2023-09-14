package controllers

import "github.com/omid-h70/bookstore/users-api/models/db"

type Handlers struct {
	store *db.Store
}

func NewHandler(store *db.Store) Handlers {
	return Handlers{store}
}
