package user

import (
	"context"

	db "github.com/M-Xue/go-auth-server/db/sqlc"
	"github.com/M-Xue/go-auth-server/server"
)

func SignUp(
	server server.Server,
	firstName string,
	lastName string,
	email string,
	username string,
	password string,
) (*db.CreateUserRow, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	createUserArg := db.CreateUserParams{
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		FirstName: firstName,
		LastName:  lastName,
	}
	newUser, err := server.DbStore.CreateUser(context.Background(), createUserArg)

	return &newUser, err
}

func LogIn(server server.Server, email string, password string) (*db.GetUserByEmailRow, error) {
	user, err := server.DbStore.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	if checkPasswordHash(password, user.Password) {
		return &user, nil
	} else {
		return nil, &InvalidCredentialsError{}
	}
}
