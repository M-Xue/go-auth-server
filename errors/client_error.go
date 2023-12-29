package errors

// ************************************************************************************************
// * Server Request Errors
// ************************************************************************************************

type ClientErrorResponse struct {
	Error         error           `json:"-"`
	ClientMessage string          `json:"message"`
	StatusCode    int             `json:"-"`
	ErrorCode     ServerErrorCode `json:"errorCode"`
	DebugId       string          `json:"debugId"`
}

func CreateClientErrorResponse(
	err error,
	msg string,
	statusCode int,
	errorCode ServerErrorCode,
	debugId string,
) ClientErrorResponse {
	return ClientErrorResponse{
		Error:         err,
		ClientMessage: msg,
		StatusCode:    statusCode,
		ErrorCode:     errorCode,
		DebugId:       debugId,
	}
}

type ClientError interface {
	GetClientErrorResponse() ClientErrorResponse
}
