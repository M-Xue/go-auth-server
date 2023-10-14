package server

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/M-Xue/go-auth-server/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	DbConn           *sql.DB
	Router           *gin.Engine
	AuthTokenFactory token.Factory
	Config           *Config
}

func InitServer() (Server, error) {
	serverConfig, err := loadConfig()
	if err != nil {
		return Server{}, err
	}

	dbURL := serverConfig.DbUrl
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	authTokenFactory, err := token.NewPasetoFactory(serverConfig.TokenSymmetricKey)
	if err != nil {
		return Server{}, fmt.Errorf("cannot create auth token factory: %w", err)
	}

	server := Server{
		DbConn:           db,
		Router:           gin.Default(),
		AuthTokenFactory: authTokenFactory,
		Config:           serverConfig,
	}

	return server, nil
}
