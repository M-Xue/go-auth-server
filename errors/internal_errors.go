package errors

import (
	"net/http"

	goerr "github.com/go-errors/errors"
	"github.com/google/uuid"
)

type InternalError struct {
	ServerError
}

func NewInternalServiceError(err error) InternalError {
	errMetadata := ErrorMetadata{
		ErrorCode: InternalServiceError,
		DebugId:   uuid.New(),
	}

	errClientRes := ErrorClientResponseData{
		StatusCode:    http.StatusInternalServerError,
		ClientMessage: "Something went wrong on the server",
	}

	serverErr := ServerError{
		ErrorMetadata:           errMetadata,
		ErrorClientResponseData: errClientRes,
		StackTraceError:         goerr.WrapPrefix(err, "Internal service error:", 0),
	}

	return InternalError{serverErr}
}
