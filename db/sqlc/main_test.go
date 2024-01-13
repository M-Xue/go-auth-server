package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(
		context.Background(),
		"postgres://maxxue@localhost:5432/go-server-template",
	)
	if err != nil {
		log.Fatal("could not connect to database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
