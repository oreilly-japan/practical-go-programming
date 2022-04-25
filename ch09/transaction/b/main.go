package main

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// transaction-defer-start
type Service struct {
	db *sql.DB
}

func (s *Service) UpdateProduct(ctx context.Context, productID string) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err = tx.ExecContext(ctx, "...", productID); err != nil {
		return err
	}

	return tx.Commit()
}

// transaction-defer-end
