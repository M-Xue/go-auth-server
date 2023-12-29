package user

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/M-Xue/go-auth-server/errors"
)

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}

func (e *UserNotFoundError) GetClientErrorResponse() errors.ClientErrorResponse {
	return errors.CreateClientErrorResponse(
		e,
		e.Error(),
		http.StatusNotFound,
		errors.UserNotFound,
		uuid.New().String(),
	)
}

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "invalid credentials"
}

func (e *InvalidCredentialsError) GetClientErrorResponse() errors.ClientErrorResponse {
	return errors.CreateClientErrorResponse(
		e,
		e.Error(),
		http.StatusUnauthorized,
		errors.InvalidCredentials,
		uuid.New().String(),
	)
}

type ExistingEmailError struct{}

func (e *ExistingEmailError) Error() string {
	return "email already exists"
}

func (e *ExistingEmailError) GetClientErrorResponse() errors.ClientErrorResponse {
	return errors.CreateClientErrorResponse(
		e,
		e.Error(),
		http.StatusConflict,
		errors.ExistingEmail,
		uuid.New().String(),
	)
}

type ExistingUsernameError struct{}

func (e *ExistingUsernameError) Error() string {
	return "username already exists"
}

func (e *ExistingUsernameError) GetClientErrorResponse() errors.ClientErrorResponse {
	return errors.CreateClientErrorResponse(
		e,
		e.Error(),
		http.StatusConflict,
		errors.ExistingUsername,
		uuid.New().String(),
	)
}
