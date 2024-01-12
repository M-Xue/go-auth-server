package api

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/M-Xue/go-auth-server/db"
	sqlc "github.com/M-Xue/go-auth-server/db/sqlc"
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
) (*sqlc.CreateUserRow, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, errors.NewInternalServiceError(err, zerolog.ErrorLevel)
	}

	createUserArg := sqlc.CreateUserParams{
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		FirstName: firstName,
		LastName:  lastName,
	}
	newUser, err := server.DbStore.CreateUser(context.Background(), createUserArg)
	if err != nil {
		if db.GetDbErrorCode(err) == db.UniqueViolation {
			// TODO figureout what the message is
			return nil, errors.NewExistingUsernameError()
		}
		return nil, errors.NewInternalServiceError(err, zerolog.ErrorLevel)
	}

	return &newUser, nil
}

func LogIn(server server.Server, email string, password string) (*sqlc.GetUserByEmailRow, error) {
	user, err := server.DbStore.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, errors.NewInternalServiceError(err, zerolog.ErrorLevel)
	}
	// TODO need to check if nil is returned when the user does not exist or if pgx error is returned

	if !isPasswordEqualToHashedPassword(password, user.Password) {
		return nil, errors.NewInvalidCredentialsError()
	}

	return &user, nil
}
