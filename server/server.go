package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	DbConn *sql.DB
	Router *gin.Engine
}
