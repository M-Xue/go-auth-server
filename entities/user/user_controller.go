package user

import (
	"context"

	db "github.com/M-Xue/go-auth-server/db/sqlc"
	"github.com/M-Xue/go-auth-server/errors"
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
		return nil, errors.NewInternalServiceError(err)
	}

	createUserArg := db.CreateUserParams{
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		FirstName: firstName,
		LastName:  lastName,
	}
	newUser, err := server.DbStore.CreateUser(context.Background(), createUserArg)
	if err != nil {
		// TODO check for existing user, conflicting credentials etc
		return nil, errors.NewInternalServiceError(err)
	}

	return &newUser, nil
}

func LogIn(server server.Server, email string, password string) (*db.GetUserByEmailRow, error) {
	user, err := server.DbStore.GetUserByEmail(context.Background(), email)
	if err != nil {
		// TODO check if user email exists
		return nil, errors.NewInternalServiceError(err)
	}
	if isPasswordEqualToHashedPassword(password, user.Password) {
		return &user, nil
	} else {
		return nil, errors.NewInvalidCredentialsError()
	}
}
