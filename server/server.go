package server

import (
	"github.com/gin-gonic/gin"

	"github.com/M-Xue/go-auth-server/auth"
	db "github.com/M-Xue/go-auth-server/db/sqlc"
)

type Server struct {
	DbStore          *db.Store
	Router           *gin.Engine
	Config           Config
	AuthTokenFactory auth.TokenFactory
}

func InitServer() (Server, error) {
	serverConfig, err := initConfig()
	if err != nil {
		return Server{}, err
	}

	dbStore, err := initDatabase(serverConfig.DbUrl)
	if err != nil {
		return Server{}, err
	}

	tokenFactory, err := auth.NewPasetoFactory(serverConfig.TokenSymmetricKey)
	if err != nil {
		return Server{}, err
	}

	server := Server{
		DbStore:          dbStore,
		Router:           gin.Default(),
		Config:           serverConfig,
		AuthTokenFactory: tokenFactory,
	}

	return server, nil
}
