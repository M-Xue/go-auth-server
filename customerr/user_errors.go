package customerr

import (
	"net/http"

	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
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
		LogLevel:        zerolog.InfoLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: errors.Errorf("user not found"),
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
		LogLevel:        zerolog.InfoLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: errors.Errorf("invalid credentials"),
	}

	return InvalidCredentialsError{serverErr}
}

type EmailNotFoundError struct {
	ServerError
}

func NewEmailNotFoundError() EmailNotFoundError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: InvalidCredentials,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusUnauthorized,
		ClientMessage:   "Invalid credentials",
		LogLevel:        zerolog.InfoLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: errors.Errorf("email not found"),
	}

	return EmailNotFoundError{serverErr}
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
		LogLevel:        zerolog.InfoLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: errors.Errorf("email already exists"),
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
		LogLevel:        zerolog.InfoLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: errors.Errorf("username already exists"),
	}

	return ExistingUsernameError{serverErr}
}
