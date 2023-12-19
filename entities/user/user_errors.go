package user

import (
	"net/http"

	"github.com/google/uuid"

	clienterror "github.com/M-Xue/go-auth-server/clienterror"
)

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}

func (e *UserNotFoundError) GetClientErrorResponse() clienterror.ClientErrorResponse {
	return clienterror.CreateClientErrorResponse(
		e,
		e.Error(),
		http.StatusNotFound,
		clienterror.UserNotFoundErrorCode,
		uuid.New().String(),
	)
}

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "invalid credentials"
}

func (e *InvalidCredentialsError) GetClientErrorResponse() clienterror.ClientErrorResponse {
	return clienterror.CreateClientErrorResponse(
		e,
		e.Error(),
		http.StatusUnauthorized,
		clienterror.InvalidCredentialsErrorCode,
		uuid.New().String(),
	)
}

type ExistingEmailError struct{}

func (e *ExistingEmailError) Error() string {
	return "email already exists"
}

func (e *ExistingEmailError) GetClientErrorResponse() clienterror.ClientErrorResponse {
	return clienterror.CreateClientErrorResponse(
		e,
		e.Error(),
		http.StatusConflict,
		clienterror.ExistingEmailErrorCode,
		uuid.New().String(),
	)
}

type ExistingUsernameError struct{}

func (e *ExistingUsernameError) Error() string {
	return "username already exists"
}

func (e *ExistingUsernameError) GetClientErrorResponse() clienterror.ClientErrorResponse {
	return clienterror.CreateClientErrorResponse(
		e,
		e.Error(),
		http.StatusConflict,
		clienterror.ExistingUsernameErrorCode,
		uuid.New().String(),
	)
}
