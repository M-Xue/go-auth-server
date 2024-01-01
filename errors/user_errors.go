package errors

import (
	"net/http"

	goerr "github.com/go-errors/errors"
	"github.com/google/uuid"
)

type UserNotFoundError struct {
	ServerError
}

func NewUserNotFoundError() UserNotFoundError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: UserNotFound,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusNotFound,
		ClientMessage:   "User not found",
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.Errorf("user not found"),
	}

	return UserNotFoundError{serverErr}
}

type InvalidCredentialsError struct {
	ServerError
}

func NewInvalidCredentialsError() InvalidCredentialsError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: InvalidCredentials,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusUnauthorized,
		ClientMessage:   "Invalid credentials",
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.Errorf("invalid credentials"),
	}

	return InvalidCredentialsError{serverErr}
}

type ExistingEmailError struct {
	ServerError
}

func NewExistingEmailError() ExistingEmailError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: ExistingEmail,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusConflict,
		ClientMessage:   "Email already exists",
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.Errorf("email already exists"),
	}

	return ExistingEmailError{serverErr}
}

type ExistingUsernameError struct {
	ServerError
}

func NewExistingUsernameError() ExistingUsernameError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: ExistingUsername,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusConflict,
		ClientMessage:   "Username already exists",
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.Errorf("username already exists"),
	}

	return ExistingUsernameError{serverErr}
}
