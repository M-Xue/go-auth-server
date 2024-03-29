package customerr

import (
	"net/http"

	goerr "github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type InternalError struct {
	ServerError
}

func NewInternalServiceError(err error, logLevel zerolog.Level) InternalError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: InternalServiceError,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusInternalServerError,
		ClientMessage:   "Something went wrong on the server",
		LogLevel:        logLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.WrapPrefix(err, "internal service error", 0),
	}

	return InternalError{serverErr}
}

func NewUncaughtInternalServiceError(err error, logLevel zerolog.Level) InternalError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: UncaughtInternalServiceError,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusInternalServerError,
		ClientMessage:   "Something went wrong on the server",
		LogLevel:        logLevel,
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.WrapPrefix(err, "uncaught internal service error", 0),
	}

	return InternalError{serverErr}
}
