package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/M-Xue/go-auth-server/middleware"
	"github.com/M-Xue/go-auth-server/server"
)

// Source: https://stackoverflow.com/questions/62608429/how-to-combine-group-of-routes-in-gin
func AttachRoutes(server server.Server) {
	// TODO: fix this
	server.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{server.Config.ClientUrl},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == server.Config.ClientUrl
		},
		MaxAge: 12 * time.Hour,
	}))

	server.Router.Use(middleware.SetRequestIdMiddleware())
	server.Router.Use(middleware.ErrorHandlerMiddleware(server.Logger))
	rootGroup := server.Router.Group("/api/v1")

	userRoutes(server, rootGroup)

	protectedRouter := rootGroup.Group("/").Use(middleware.AuthenticationMiddleware(server))
	protectedRouter.GET("/secret", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"secret": "Foo 42",
		})
	})
}
