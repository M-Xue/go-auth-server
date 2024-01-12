package errors

type ServerErrorCode int

const (
	UserNotFound       ServerErrorCode = 101
	InvalidCredentials ServerErrorCode = 102
	ExistingEmail      ServerErrorCode = 103
	ExistingUsername   ServerErrorCode = 104

	InternalServiceError         ServerErrorCode = 200
	UncaughtInternalServiceError ServerErrorCode = 201
)
