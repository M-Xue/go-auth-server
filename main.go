package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/M-Xue/go-auth-server/middleware"
	"github.com/M-Xue/go-auth-server/routes"
	"github.com/M-Xue/go-auth-server/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	server := server.Server{
		DbConn: db,
		Router: gin.Default(),
	}

	server.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))
	server.Router.Use(middleware.ErrorHandlerMiddleware())

	rootGroup := server.Router.Group("/api/v1")
	routes.AttachRoutes(server, rootGroup)
	server.Router.Run("localhost:8080")
}
