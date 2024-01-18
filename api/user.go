package api

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog"

	"github.com/M-Xue/go-auth-server/customerr"
	"github.com/M-Xue/go-auth-server/db"
	sqlc "github.com/M-Xue/go-auth-server/db/sqlc"
	"github.com/M-Xue/go-auth-server/server"
	"github.com/go-errors/errors"
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
		return nil, customerr.NewInternalServiceError(err, zerolog.ErrorLevel)
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
		pgerr, ok := err.(*pgconn.PgError)
		if ok && pgerr.Code == db.UniqueViolation {
			if pgerr.ConstraintName == "users_username_key" {
				return nil, customerr.NewExistingUsernameError()
			} else if pgerr.ConstraintName == "users_email_key" {
				return nil, customerr.NewExistingEmailError()
			}
		}
		return nil, customerr.NewInternalServiceError(err, zerolog.ErrorLevel)
	}

	return &newUser, nil
}

func LogIn(server server.Server, email string, password string) (*sqlc.GetUserByEmailRow, error) {
	user, err := server.DbStore.GetUserByEmail(context.Background(), email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, customerr.NewEmailNotFoundError()
		}
		return nil, customerr.NewInternalServiceError(err, zerolog.ErrorLevel)
	}

	if !isPasswordEqualToHashedPassword(password, user.Password) {
		return nil, customerr.NewInvalidCredentialsError()
	}

	return &user, nil
}
