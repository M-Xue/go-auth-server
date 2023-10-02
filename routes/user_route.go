package routes

import (
	"net/http"

	"github.com/M-Xue/go-auth-server/entities/user"
	"github.com/M-Xue/go-auth-server/server"
	"github.com/gin-gonic/gin"
)

func userRoutes(server server.Server, rootGroup *gin.RouterGroup) {
	userRouter := rootGroup.Group("/user")
	{
		userRouter.POST("/", handleUserSignUp(server))
		userRouter.POST("/auth/email", handleIsEmailAvailable(server))
		userRouter.POST("/auth/username", handleIsUsernameAvailable(server))
		userRouter.POST("/auth/login", handleUserLogin(server))
	}
}

// ************************************************************************************************
// * Handlers
// ************************************************************************************************

func handleUserSignUp(server server.Server) gin.HandlerFunc {
	type request struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}
	type response struct {
		ID        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Username  string `json:"username"`
	}
	handler := func(ctx *gin.Context) {
		var req request
		if err := ctx.BindJSON(&req); err != nil {
			return
		}

		newUser, err := user.UserSignUp(
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

		err = setJwtAuthenticationCookie(ctx, newUser)
		if err != nil {
			ctx.Error(err)
			return
		}

		res := response{
			ID:        newUser.ID,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Email:     newUser.Email,
			Username:  newUser.Username,
		}
		ctx.JSON(http.StatusOK, res)
	}
	return handler
}

func handleUserLogin(server server.Server) gin.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		ID        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Username  string `json:"username"`
	}
	handler := func(ctx *gin.Context) {
		var req request
		if err := ctx.BindJSON(&req); err != nil {
			return
		}
		loggedInUser, err := user.UserLogin(server, req.Email, req.Password)
		if err != nil {
			ctx.Error(err)
			return
		}

		err = setJwtAuthenticationCookie(ctx, loggedInUser)
		if err != nil {
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

func handleIsEmailAvailable(server server.Server) gin.HandlerFunc {
	type request struct {
		Email string `json:"email"`
	}
	type response struct {
		IsEmailAvailable bool `json:"isEmailAvailable"`
	}
	handler := func(ctx *gin.Context) {
		var req request
		if err := ctx.BindJSON(&req); err != nil {
			return
		}

		isEmailAvail, err := user.UserIsEmailAvailable(server, req.Email)
		if err != nil {
			ctx.Error(err)
			return
		}

		res := response{
			IsEmailAvailable: isEmailAvail,
		}
		ctx.JSON(http.StatusOK, res)
	}
	return handler
}

func handleIsUsernameAvailable(server server.Server) gin.HandlerFunc {
	type request struct {
		Username string `json:"username"`
	}
	type response struct {
		IsUsernameAvailable bool `json:"isUsernameAvailable"`
	}
	handler := func(ctx *gin.Context) {
		var req request
		if err := ctx.BindJSON(&req); err != nil {
			return
		}

		isUsernameAvail, err := user.UserIsUsernameAvailable(server, req.Username)
		if err != nil {
			ctx.Error(err)
			return
		}

		res := response{
			IsUsernameAvailable: isUsernameAvail,
		}
		ctx.JSON(http.StatusOK, res)
	}
	return handler
}
