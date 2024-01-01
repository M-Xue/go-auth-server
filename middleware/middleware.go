package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/M-Xue/go-auth-server/errors"
	"github.com/M-Xue/go-auth-server/server"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case errors.InternalError:
				// TODO add logging here
				c.AbortWithStatusJSON(e.GetClientErrorResponse().HttpStatusCode, e)
			default:
				uuid := uuid.New().String()
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"debugId": uuid})
			}
		}
	}
}

// // TODO: check the tutorials errorResponse() function
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
		c.Set("user", user)
		c.Next()
	}
}
