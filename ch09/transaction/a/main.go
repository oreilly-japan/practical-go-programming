package main

import (
	"context"
	"database/sql"
)

// transaction-basic-start
type Service struct {
	db *sql.DB
}

func (s *Service) UpdateProduct(ctx context.Context, productID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, "UPDATE products SET price = 200 WHERE product_id = $1", productID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// transaction-basic-end
