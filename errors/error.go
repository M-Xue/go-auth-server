package errors

import (
	goerr "github.com/go-errors/errors"
	"github.com/google/uuid"
)

type ErrorMetadata struct {
	ErrorCode ServerErrorCode
	DebugId   uuid.UUID
}

type ErrorClientResponseData struct {
	StatusCode    int
	ClientMessage string
}

type StackTraceError = *goerr.Error

type ServerError struct {
	ErrorMetadata
	ErrorClientResponseData
	StackTraceError // Use goerror.Wrap is there is underlying error, else use goerr.Errorf
}

type ClientErrorResponse struct {
	DebugId       uuid.UUID `json:"debug_id"`
	StatusCode    int       `json:"-"`
	ClientMessage string    `json:"message"`
}

func (err ServerError) GetClientErrorResponse() ClientErrorResponse {
	return ClientErrorResponse{
		DebugId:       err.DebugId,
		StatusCode:    err.StatusCode,
		ClientMessage: err.ClientMessage,
	}
}

// what do i want
// i want all my errors to inherit from the go-error class

// I want it to already have the stack trace attached to it
// i want to be able to use as and is
// i want it to return a client error response
// I want it to have internal server metadata like internal server error codes

// I want structs holding the servererror to be errors themselves so i can use Is on them
// I want to be able to use Is on nested errors within these

// i think i need a new type for this??

// inheritance of classes (structs)
// interface inheritance
// default methods
// multi inheritance
