package routes

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/M-Xue/go-auth-server/entities/user"
	"github.com/M-Xue/go-auth-server/server"
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

func setPasetoAuthCookie(ctx *gin.Context, user *user.User, server server.Server) error {
	tokenDuration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		return err
	}
	accessToken, err := server.AuthTokenFactory.CreateAuthToken(user.Username, tokenDuration)

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("auth", accessToken, 9999999, "", "", false, true) // TODO: fix these args
	return nil
}
