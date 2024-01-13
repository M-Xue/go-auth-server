package util

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/M-Xue/go-auth-server/errors"
)

func LogError(
	c *gin.Context,
	logger zerolog.Logger,
	e errors.ServerError,
) {
	userIDString := ""
	userUUID, err := GetUserUUIDFromGinContext(c)
	if err != nil {
		userIDString = userUUID.String()
	}

	requestIDString := ""
	requestUUID, err := GetRequestUUIDFromGinContext(c)
	if err != nil {
		requestIDString = requestUUID.String()
	}

	requestInfo := GetRequestInfoFromGinContext(c)
	requestInfoString, err := json.Marshal(requestInfo)
	if err != nil {
		// TODO
		return
	}

	logger.
		WithLevel(e.LogLevel).
		Str("debug_id", e.DebugId.String()).
		Str("user_id", userIDString).
		Str("request_id", requestIDString).
		Str("request_info", string(requestInfoString)).
		Err(e).
		Str("stack_trace", string(e.Stack())).
		Send()
}
