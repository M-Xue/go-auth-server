package customerr

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	goerr "github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type fooError struct {
	ServerError
}

func NewFooError() fooError {
	errMetadata := ErrorMetadata{
		DebugId:         uuid.New(),
		ServerErrorCode: InternalServiceError,
		HttpStatusCode:  http.StatusInternalServerError,
		ClientMessage:   "Something went wrong on the server",
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.Errorf("test message"),
	}

	return fooError{serverErr}
}

func NewFooErrorWithNestedError(err error) fooError {
	errMetadata := ErrorMetadata{
		ServerErrorCode: InternalServiceError,
		DebugId:         uuid.New(),
		HttpStatusCode:  http.StatusInternalServerError,
		ClientMessage:   "Test ClientMessage",
	}

	serverErr := ServerError{
		ErrorMetadata:   errMetadata,
		StackTraceError: goerr.WrapPrefix(err, "test error:", 0),
	}

	return fooError{serverErr}
}

// TESTED BEHAVIOURS
// Custom error structs implement the error interface (behaves like error)
// Can wrap them in other errors
// Can use As and Is from standard library
// Can unwrap them using standard library
// Custom error structs have stack traces attached
// Check client error response

func TestStandardLibErrorsIsOnServerError(t *testing.T) {
	fooErr := NewFooError()
	wrappedErr := fmt.Errorf("wrapping error: %w", fooErr)
	require.True(t, errors.Is(wrappedErr, ServerError{}))
}

func TestStandardLibErrorsAsOnServerError(t *testing.T) {
	fooErr := NewFooError()
	wrappedErr := fmt.Errorf("wrapping error: %w", fooErr)

	assignedErr := &ServerError{}
	require.True(t, errors.As(wrappedErr, assignedErr))
	require.Equal(t, fooErr.DebugId, assignedErr.DebugId)
	require.Equal(t, fooErr.ServerErrorCode, assignedErr.ServerErrorCode)
	require.Equal(t, fooErr.HttpStatusCode, assignedErr.HttpStatusCode)
	require.Equal(t, fooErr.ClientMessage, assignedErr.ClientMessage)
}

func TestStandardLibErrorsUnwrapOnServerError(t *testing.T) {
	fooErr := NewFooError()
	wrappedErr := fmt.Errorf("%w", fooErr)

	unwrappedErr := errors.Unwrap(wrappedErr)
	convertedFooErr, ok := unwrappedErr.(fooError)
	require.True(t, ok)
	if ok {
		require.Equal(t, fooErr.DebugId, convertedFooErr.DebugId)
		require.Equal(t, fooErr.ServerErrorCode, convertedFooErr.ServerErrorCode)
		require.Equal(t, fooErr.HttpStatusCode, convertedFooErr.HttpStatusCode)
		require.Equal(t, fooErr.ClientMessage, convertedFooErr.ClientMessage)
	}
}

func TestServerErrorStackTrace(t *testing.T) {
	fooErr := NewFooError()
	stackTrace := fooErr.Stack()
	require.NotNil(t, stackTrace)
	errorStack := fooErr.ErrorStack()
	require.NotNil(t, errorStack)
}

func TestServerErrorClientResponse(t *testing.T) {
	fooErr := NewFooError()
	clientRes := fooErr.GetClientErrorResponse()
	require.Equal(t, fooErr.DebugId, clientRes.DebugId)
	require.Equal(t, fooErr.ServerErrorCode, clientRes.ServerErrorCode)
	require.Equal(t, fooErr.HttpStatusCode, clientRes.HttpStatusCode)
	require.Equal(t, fooErr.ClientMessage, clientRes.ClientMessage)
}
