package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/M-Xue/go-auth-server/errors"
	"github.com/M-Xue/go-auth-server/server"
	"github.com/M-Xue/go-auth-server/util"
	goerr "github.com/go-errors/errors"
)

func ErrorHandlerMiddleware(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			e := &errors.ServerError{}
			if goerr.As(err, e) {
				util.LogError(c, logger, *e)
				c.AbortWithStatusJSON(
					e.GetClientErrorResponse().HttpStatusCode,
					e.GetClientErrorResponse(),
				)
			} else {
				uncaughtError := errors.NewUncaughtInternalServiceError(err, zerolog.ErrorLevel)
				util.LogError(c, logger, uncaughtError.ServerError)
				c.AbortWithStatusJSON(uncaughtError.GetClientErrorResponse().HttpStatusCode, uncaughtError.GetClientErrorResponse())
			}

		}
	}
}

func SetRequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.New()
		c.Set(util.RequestIDGinContextKey, requestId)
		c.Next()
	}
}

func AuthenticationMiddleware(server server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("auth")
		if err != nil {
			e := errors.NewMissingAuthTokenError()
			c.AbortWithStatusJSON(
				e.GetClientErrorResponse().HttpStatusCode,
				e.GetClientErrorResponse(),
			)
			return
		}

		payload, err := server.AuthTokenFactory.VerifyAndParseAuthToken(tokenString)
		if err != nil {
			e := errors.NewInvalidAuthTokenError()
			c.AbortWithStatusJSON(
				e.GetClientErrorResponse().HttpStatusCode,
				e.GetClientErrorResponse(),
			)
			return
		}

		userIdClaim := payload.TokenID
		c.Set(util.UserIDGinContextKey, userIdClaim)
		c.Next()
	}
}
