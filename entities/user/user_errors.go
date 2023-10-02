package user

import (
	"net/http"

	"github.com/M-Xue/go-auth-server/customerror"
	"github.com/google/uuid"
)

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}

func (e *UserNotFoundError) GetClientErrorResponse() customerror.ClientErrorResponse {
	return customerror.CreateClientErrorResponse(e, e.Error(), http.StatusNotFound, customerror.UserNotFoundErrorCode, uuid.New().String())
}

type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "invalid credentials"
}

func (e *InvalidCredentialsError) GetClientErrorResponse() customerror.ClientErrorResponse {
	return customerror.CreateClientErrorResponse(e, e.Error(), http.StatusUnauthorized, customerror.InvalidCredentialsErrorCode, uuid.New().String())
}

type ExistingEmailError struct{}

func (e *ExistingEmailError) Error() string {
	return "email already exists"
}

func (e *ExistingEmailError) GetClientErrorResponse() customerror.ClientErrorResponse {
	return customerror.CreateClientErrorResponse(e, e.Error(), http.StatusConflict, customerror.ExistingEmailErrorCode, uuid.New().String())
}

type ExistingUsernameError struct{}

func (e *ExistingUsernameError) Error() string {
	return "username already exists"
}

func (e *ExistingUsernameError) GetClientErrorResponse() customerror.ClientErrorResponse {
	return customerror.CreateClientErrorResponse(e, e.Error(), http.StatusConflict, customerror.ExistingUsernameErrorCode, uuid.New().String())
}
