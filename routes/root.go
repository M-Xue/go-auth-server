package routes

import (
	"net/http"

	"github.com/M-Xue/go-auth-server/middleware"
	"github.com/M-Xue/go-auth-server/server"
	"github.com/gin-gonic/gin"
)

// Source: https://stackoverflow.com/questions/62608429/how-to-combine-group-of-routes-in-gin
func AttachRoutes(server server.Server, rootGroup *gin.RouterGroup) {
	userRoutes(server, rootGroup)

	protectedRouter := rootGroup.Group("/protected")
	protectedRouter.Use(middleware.AuthenticationMiddleware(server))
	protectedRouter.GET("/secret", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"secret": "Foo 42",
		})
	})
}
