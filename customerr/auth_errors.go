package customerr

import (
	"net/http"

	goerr "github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type MissingAuthTokenError struct {
	ServerError
}

func NewMissingAuthTokenError() MissingAuthTokenError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: UncaughtInternalServiceError,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusUnauthorized,
		ClientMessage:   "Missing authentication token",
		LogLevel:        zerolog.InfoLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.Errorf("missing authentication token"),
	}

	return MissingAuthTokenError{serverErr}
}

type InvalidAuthTokenError struct {
	ServerError
}

func NewInvalidAuthTokenError() InvalidAuthTokenError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: UncaughtInternalServiceError,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusUnauthorized,
		ClientMessage:   "Invalid authentication token",
		LogLevel:        zerolog.InfoLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.Errorf("authentication token is invalid"),
	}

	return InvalidAuthTokenError{serverErr}
}

type ExpiredAuthTokenError struct {
	ServerError
}

func NewExpiredAuthTokenError() ExpiredAuthTokenError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: UncaughtInternalServiceError,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusUnauthorized,
		ClientMessage:   "Expired authentication token",
		LogLevel:        zerolog.InfoLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.Errorf("authentication token has expired"),
	}

	return ExpiredAuthTokenError{serverErr}
}
