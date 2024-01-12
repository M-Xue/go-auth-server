package api

import (
	"log"
	"os"
	"testing"

	"github.com/M-Xue/go-auth-server/server"
)

var testServer server.Server

func TestMain(m *testing.M) {
	server, err := server.InitServer()
	if err != nil {
		log.Fatal("could not initialize server: %w", err.Error())
	}

	testServer = server

	os.Exit(m.Run())
}
