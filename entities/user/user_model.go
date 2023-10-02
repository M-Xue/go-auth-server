package user

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/M-Xue/go-auth-server/server"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
}

func createUser(server server.Server, firstName string, lastName string, email string, username string, password string) (*User, error) {
	query := "INSERT INTO `user` (`id`, `firstName`, `lastName`, `email`, `username`, `password`) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := server.DbConn.Query(query, uuid.New().String(), firstName, lastName, email, username, password)
	// TODO: fix this underscore and get back the user

	if err != nil {
		var driverErr *mysql.MySQLError
		if errors.As(err, &driverErr) {
			if driverErr.Number == 1062 {
				if strings.Contains(err.Error(), "user.email") {
					return nil, &ExistingEmailError{}
				}
				if strings.Contains(err.Error(), "user.username") {
					return nil, &ExistingUsernameError{}
				}
			}
		}
		return nil, err
	}

	return &User{ID: uuid.New().String(), FirstName: firstName, LastName: lastName, Email: email, Username: username, Password: ""}, nil
}

func getUserById(server server.Server, id string) (*User, error) {
	var user User
	query := "SELECT id, firstName, lastName, email, username, password FROM `user` WHERE `id` = ?"
	err := server.DbConn.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &UserNotFoundError{}
		}
		return nil, err
	}
	return &user, nil
}

func getUserByEmail(server server.Server, email string) (*User, error) {
	var user User
	query := "SELECT id, firstName, lastName, email, username, password FROM `user` WHERE `email` = ?"
	err := server.DbConn.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &UserNotFoundError{}
		}
		return nil, err
	}
	return &user, nil
}

func getUserByUsername(server server.Server, username string) (*User, error) {
	var user User
	query := "SELECT id, firstName, lastName, email, username, password FROM `user` WHERE `username` = ?"
	err := server.DbConn.QueryRow(query, username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &UserNotFoundError{}
		}
		return nil, err
	}
	return &user, nil
}
