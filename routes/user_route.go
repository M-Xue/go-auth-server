package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/M-Xue/go-auth-server/entities/user"
	"github.com/M-Xue/go-auth-server/server"
)

func userRoutes(server server.Server, rootGroup *gin.RouterGroup) {
	userRouter := rootGroup.Group("/user")
	{
		userRouter.POST("/auth/signup", handleUserSignUp(server))
		userRouter.POST("/auth/login", handleUserLogIn(server))
	}
}

// ************************************************************************************************
// * Handlers
// ************************************************************************************************

func handleUserSignUp(server server.Server) gin.HandlerFunc {
	type request struct {
		Email     string `json:"email"`
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Password  string `json:"password"`
	}
	type response struct {
		ID        uuid.UUID `json:"id"`
		Email     string    `json:"email"`
		Username  string    `json:"username"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
	}
	handler := func(ctx *gin.Context) {
		var req request
		if err := ctx.BindJSON(&req); err != nil {
			return
		}

		newUser, err := user.SignUp(
			server,
			req.FirstName,
			req.LastName,
			req.Email,
			req.Username,
			req.Password,
		)
		if err != nil {
			ctx.Error(err)
			return
		}

		err = setAuthTokenCookie(ctx, server, newUser.ID)
		if err != nil {
			ctx.Error(err)
			return
		}

		res := response{
			ID:        newUser.ID,
			Email:     newUser.Email,
			Username:  newUser.Username,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
		}
		ctx.JSON(http.StatusOK, res)
	}
	return handler
}

func handleUserLogIn(server server.Server) gin.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		ID        uuid.UUID `json:"id"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Email     string    `json:"email"`
		Username  string    `json:"username"`
	}
	handler := func(ctx *gin.Context) {
		var req request
		if err := ctx.BindJSON(&req); err != nil {
			return
		}
		loggedInUser, err := user.LogIn(server, req.Email, req.Password)
		if err != nil {
			ctx.Error(err)
			return
		}

		err = setAuthTokenCookie(ctx, server, loggedInUser.ID)
		if err != nil {
			// TODO fix if this error is the correct way to throw this err
			ctx.Error(err)
			return
		}

		res := response{
			ID:        loggedInUser.ID,
			FirstName: loggedInUser.FirstName,
			LastName:  loggedInUser.LastName,
			Email:     loggedInUser.Email,
			Username:  loggedInUser.Username,
		}
		ctx.JSON(http.StatusOK, res)
	}
	return handler
}
