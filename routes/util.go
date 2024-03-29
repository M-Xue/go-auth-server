package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/M-Xue/go-auth-server/customerr"
	"github.com/M-Xue/go-auth-server/server"
)

func setAuthTokenCookie(ctx *gin.Context, server server.Server, userID uuid.UUID) error {
	tokenDuration, err := time.ParseDuration(server.Config.AccessDuration)
	if err != nil {
		return customerr.NewInternalServiceError(err, zerolog.ErrorLevel)
	}
	accessToken, err := server.AuthTokenFactory.CreateAuthToken(userID, tokenDuration)
	if err != nil {
		return customerr.NewInternalServiceError(err, zerolog.ErrorLevel)
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("auth", accessToken, 9999999, "", "", false, true) // TODO: fix these args
	return nil
}
