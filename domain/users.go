package domain

import (
	"errors"
	"github.com/omid-h70/bookstore/users-api/models/db"
	"strings"
)

type User struct {
	UserId    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func Validate(u db.CreateUserParam) error {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.New("invalid email address")
	}
	return nil
}
