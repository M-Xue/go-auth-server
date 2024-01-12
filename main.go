package main

import (
	"github.com/M-Xue/go-auth-server/routes"
	"github.com/M-Xue/go-auth-server/server"
)

func main() {
	server, err := server.InitServer()
	if err != nil {
		panic(err)
	}
	routes.AttachRoutes(server)
	err = server.Router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
