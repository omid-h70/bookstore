package mock

import (
	"errors"
	"github.com/omid-h70/bookstore/users-api/domain"
	"github.com/omid-h70/bookstore/users-api/models/db"
)

type UserMockDB struct {
	store map[int64]domain.User
}

func NewUserMockDB() UserMockDB {
	return UserMockDB{
		store: make(map[int64]domain.User),
	}
}

func (u *UserMockDB) GetUser(userId int64) (domain.User, error) {
	user, ok := u.store[userId]
	if !ok {
		return domain.User{}, errors.New("not found")
	}
	return user, nil
}

func (u *UserMockDB) GetUserByEmail(email string) (domain.User, error) {
	for _, user := range u.store {
		if user.Email == email {
			return user, nil
		}
	}
	return domain.User{}, errors.New("not found")
}

func (u *UserMockDB) CreateUser(param db.CreateUserParam) (domain.User, error) {

	if user, err := u.GetUser(param.UserID); err != nil {
		return user, errors.New("already exist")
	}

	newUser := domain.User{
		UserId: param.UserID,
	}
	u.store[param.UserID] = newUser
	return newUser, nil
}
