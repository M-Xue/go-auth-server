package util

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/M-Xue/go-auth-server/customerr"
)

func LogError(
	c *gin.Context,
	logger zerolog.Logger,
	e customerr.ServerError,
) {
	var userIDString string
	userUUID, err := GetUserUUIDFromGinContext(c)
	if err != nil {
		userIDString = ""
	} else {
		userIDString = userUUID.String()
	}

	var requestIDString string
	requestUUID, err := GetRequestUUIDFromGinContext(c)
	if err != nil {
		requestIDString = ""
	} else {
		requestIDString = requestUUID.String()
	}

	requestInfo := GetRequestInfoFromGinContext(c)
	requestInfoJson, err := json.Marshal(requestInfo)
	var requestInfoString string
	if err != nil {
		requestInfoString = ""
	} else {
		requestInfoString = string(requestInfoJson)
	}

	logger.
		WithLevel(e.LogLevel).
		Str("debug_id", e.DebugId.String()).
		Str("user_id", userIDString).
		Str("request_id", requestIDString).
		Str("request_info", requestInfoString).
		Err(e).
		Str("stack_trace", e.ErrorStack()).
		Send()
}
