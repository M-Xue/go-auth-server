package errors

import (
	goerr "github.com/go-errors/errors"
	"github.com/google/uuid"
)

type ErrorMetadata struct {
	DebugId         uuid.UUID
	ServerErrorCode ServerErrorCode
	HttpStatusCode  int
	ClientMessage   string
}

type StackTraceError = *goerr.Error

type ServerError struct {
	ErrorMetadata
	StackTraceError // Use goerror.Wrap if there is underlying error, else use goerr.Errorf
}

type ClientErrorResponse struct {
	DebugId         uuid.UUID       `json:"debug_id"`
	ServerErrorCode ServerErrorCode `json:"error_code`
	HttpStatusCode  int             `json:"-"`
	ClientMessage   string          `json:"message"`
}

func (err ServerError) GetClientErrorResponse() ClientErrorResponse {
	return ClientErrorResponse{
		DebugId:         err.DebugId,
		ServerErrorCode: err.ServerErrorCode,
		HttpStatusCode:  err.HttpStatusCode,
		ClientMessage:   err.ClientMessage,
	}
}

func (err ServerError) Is(target error) bool {
	_, ok := target.(ServerError)
	return ok
}

func (err ServerError) As(target interface{}) bool {
	serverError, ok := target.(*ServerError)
	if ok {
		*serverError = err
	}
	return ok
}
