package server

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	db "github.com/M-Xue/go-auth-server/db/sqlc"
)

func initDatabase(dbURL string) (*db.Store, error) {
	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}
	defer dbpool.Close()

	return db.NewStore(dbpool), nil
}
