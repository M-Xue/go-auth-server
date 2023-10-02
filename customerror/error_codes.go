package customerror

type ServerErrorCode int

const (
	UserNotFoundErrorCode       ServerErrorCode = 101
	InvalidCredentialsErrorCode ServerErrorCode = 102
	ExistingEmailErrorCode      ServerErrorCode = 103
	ExistingUsernameErrorCode   ServerErrorCode = 104
)
