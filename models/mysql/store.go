package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/omid-h70/bookstore/users-api/domain"
	"github.com/omid-h70/bookstore/users-api/models/db"
	"time"
)

const (
	errMysqlDuplicateKey = 1062
	createUserQuery      = "INSERT INTO users (first_name, last_name, email, created_at) VALUES(?, ?, ?, ?);"
	getUserQuery         = "SELECT first_name, last_name, email, created_at FROM users WHERE user_id=?;"
)

type MysqlStore struct {
	client *sql.DB
}

func NewMysqlStore(dbSource string) (MysqlStore, error) {
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		return MysqlStore{}, errors.New("panic, Mysql is Down")
	}
	//you can add your logger to mysql too as well
	return MysqlStore{
		client: db,
	}, nil
}

func (m *MysqlStore) GetUser(userId int64) (domain.User, error) {
	stmnt, err := m.client.Prepare(createUserQuery)
	if err != nil {
		return domain.User{}, err
	}
	defer stmnt.Close()

	result := stmnt.QueryRow(userId)
	var user domain.User
	if err := result.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (m *MysqlStore) GetUserByEmail(email string) (domain.User, error) {
	stmnt, err := m.client.Prepare(createUserQuery)
	if err != nil {
		return domain.User{}, err
	}
	defer stmnt.Close()

	result := stmnt.QueryRow(email)
	var user domain.User
	if err := result.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (m *MysqlStore) CreateUser(param db.CreateUserParam) (domain.User, error) {
	stmnt, err := m.client.Prepare(createUserQuery)
	if err != nil {
		return domain.User{}, err
	}
	defer stmnt.Close()

	result, err := stmnt.Exec(param.FirstName, param.LastName, param.Email, param.CreatedAt)
	if err != nil {
		if tmpErr, ok := err.(*mysql.MySQLError); ok {
			if tmpErr.Number == errMysqlDuplicateKey {
				return domain.User{}, fmt.Errorf("user with %s already exist", param.Email)
			}
		}
		return domain.User{}, err
	}

	newUserId, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}

	//Default string time format
	//01-02-2006T15:04:05Z
	newUser := domain.User{
		UserId:    newUserId,
		FirstName: param.FirstName,
		LastName:  param.LastName,
		Email:     param.Email,
		CreatedAt: time.Now(),
	}

	return newUser, nil
}
