package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/M-Xue/go-auth-server/customerror"
	"github.com/M-Xue/go-auth-server/entities/user"
	"github.com/M-Xue/go-auth-server/server"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

// Source: https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac
func AuthenticationMiddleware(server server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("auth")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_HASH")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userIdClaim := claims["userId"].(string)
			user, err := user.GetUserById(server, userIdClaim)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Set("user", user)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}
