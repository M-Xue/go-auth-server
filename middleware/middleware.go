package middleware

import (
	"net/http"

	"github.com/M-Xue/go-auth-server/customerror"
	"github.com/M-Xue/go-auth-server/entities/user"
	"github.com/M-Xue/go-auth-server/server"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case customerror.ClientError:
				c.AbortWithStatusJSON(e.GetClientErrorResponse().StatusCode, e)
			default:
				uuid := uuid.New().String()
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"debugId": uuid})
			}
		}
	}
}

// TODO: check the tutorials errorResponse() function
func AuthenticationMiddleware(server server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("auth")
		if err != nil {
			// TODO: err msg: authorization header is not provided
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		payload, err := server.AuthTokenFactory.VerifyAuthToken(tokenString)
		if err != nil {
			// TODO: err msg
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userIdClaim := payload.Username

		user, err := user.GetUserById(server, userIdClaim)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
		c.Next()
	}
}
