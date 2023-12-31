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
		ErrorCode: UserNotFound,
		DebugId:   uuid.New(),
	}

	errClientRes := ErrorClientResponseData{
		StatusCode:    http.StatusNotFound,
		ClientMessage: "User not found",
	}

	serverErr := ServerError{
		ErrorMetadata:           errMetadata,
		ErrorClientResponseData: errClientRes,
		StackTraceError:         goerr.Errorf("user not found"),
	}

	return UserNotFoundError{serverErr}
}

type InvalidCredentialsError struct {
	ServerError
}

func NewInvalidCredentialsError() InvalidCredentialsError {
	errMetadata := ErrorMetadata{
		ErrorCode: InvalidCredentials,
		DebugId:   uuid.New(),
	}

	errClientRes := ErrorClientResponseData{
		StatusCode:    http.StatusUnauthorized,
		ClientMessage: "Invalid credentials",
	}

	serverErr := ServerError{
		ErrorMetadata:           errMetadata,
		ErrorClientResponseData: errClientRes,
		StackTraceError:         goerr.Errorf("invalid credentials"),
	}

	return InvalidCredentialsError{serverErr}
}

type ExistingEmailError struct {
	ServerError
}

func NewExistingEmailError() ExistingEmailError {
	errMetadata := ErrorMetadata{
		ErrorCode: ExistingEmail,
		DebugId:   uuid.New(),
	}

	errClientRes := ErrorClientResponseData{
		StatusCode:    http.StatusConflict,
		ClientMessage: "Email already exists",
	}

	serverErr := ServerError{
		ErrorMetadata:           errMetadata,
		ErrorClientResponseData: errClientRes,
		StackTraceError:         goerr.Errorf("email already exists"),
	}

	return ExistingEmailError{serverErr}
}

type ExistingUsernameError struct {
	ServerError
}

func NewExistingUsernameError() ExistingUsernameError {
	errMetadata := ErrorMetadata{
		ErrorCode: ExistingUsername,
		DebugId:   uuid.New(),
	}

	errClientRes := ErrorClientResponseData{
		StatusCode:    http.StatusConflict,
		ClientMessage: "Username already exists",
	}

	serverErr := ServerError{
		ErrorMetadata:           errMetadata,
		ErrorClientResponseData: errClientRes,
		StackTraceError:         goerr.Errorf("username already exists"),
	}

	return ExistingUsernameError{serverErr}
}
