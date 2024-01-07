package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/M-Xue/go-auth-server/errors"
	"github.com/M-Xue/go-auth-server/server"
)

const UserIDGinContextKey = "user_id"

const RequestIDGinContextKey = "request_id"

func GetUserUUIDFromGinContext(c *gin.Context) (uuid.UUID, error) {
	if userID, exists := c.Get(UserIDGinContextKey); exists {
		if userUUID, ok := userID.(uuid.UUID); ok {
			return userUUID, nil
		} else {
			return uuid.UUID{}, fmt.Errorf("user_id in gin context not uuid.UUID")
		}
	} else {
		return uuid.UUID{}, fmt.Errorf("user_id not found in gin context ")
	}
}

// TODO add request id to all requests. This is attatched to all requests but the debug id only gets added for errors

func GetRequestUUIDFromGinContext(c *gin.Context) (uuid.UUID, error) {
	if requestID, exists := c.Get(RequestIDGinContextKey); exists {
		if requestUUID, ok := requestID.(uuid.UUID); ok {
			return requestUUID, nil
		} else {
			return uuid.UUID{}, fmt.Errorf("request_id in gin context not uuid.UUID")
		}
	} else {
		return uuid.UUID{}, fmt.Errorf("request_id not found in gin context ")
	}
}

type RequestInfo struct {
	RequestUrl  string
	HttpMethod  string
	RequestBody string
	QueryParams string
}

func GetRequestInfoFromGinContext(c *gin.Context) RequestInfo {
	requestURL := c.Request.Host + c.Request.URL.Path
	httpMethod := c.Request.Method

	requestBody := ""
	requestBodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		requestBody = string(requestBodyBytes)
	}

	queryParams := c.Request.URL.Query().Encode()

	reqInfo := RequestInfo{
		RequestUrl:  requestURL,
		HttpMethod:  httpMethod,
		RequestBody: requestBody,
		QueryParams: queryParams,
	}

	return reqInfo
}

// TODO move above into server package

func LogInternalError(c *gin.Context, logger zerolog.Logger, e errors.InternalError) {
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

func ErrorHandlerMiddleware(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			e, ok := err.Err.(errors.InternalError)
			if ok {
				LogInternalError(c, logger, e)
				c.AbortWithStatusJSON(e.GetClientErrorResponse().HttpStatusCode, e)
			} else {
				uuid := uuid.New().String()
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"debugId": uuid})
			}

		}
	}
}

func AuthenticationMiddleware(server server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("auth")
		if err != nil {
			// TODO: err msg: authorization header is not provided
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		payload, err := server.AuthTokenFactory.VerifyAndParseAuthToken(tokenString)
		if err != nil {
			// TODO: err msg
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userIdClaim := payload.TokenID

		user, err := server.DbStore.GetUserByID(context.Background(), userIdClaim)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set(UserIDGinContextKey, user) // TODO turn into uuid
		c.Next()
	}
}
