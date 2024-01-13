package util

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
