package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Transations
// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(ctx); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit(ctx)
// }
