package routes

import (
	"net/http"
	"os"

	"github.com/M-Xue/go-auth-server/entities/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func setJwtAuthenticationCookie(ctx *gin.Context, user *user.User) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_HASH")))
	if err != nil {
		return err
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("auth", tokenString, 9999999, "", "", false, true) // TODO: fix these args
	return nil
}
