package main

import (
	"context"
	"database/sql"
	"fmt"
)

// transaction-wrapper-start
// txAdminはトランザクション制御するための構造体
type txAdmin struct {
	*sql.DB
}

type Service struct {
	tx txAdmin
}

// Transaction はトランザクションを制御するメソッド
// アプリケーション開発者が本メソッドを使って、DMLのクエリーを発行する
func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context) (err error)) error {
	tx, err := t.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}
	return tx.Commit()
}

func (s *Service) UpdateProduct(ctx context.Context, productID string) error {
	updateFunc := func(ctx context.Context) error {
		if _, err := s.tx.ExecContext(ctx, "..."); err != nil {
			return err
		}
		if _, err := s.tx.ExecContext(ctx, "..."); err != nil {
			return err
		}
		return nil
	}
	return s.tx.Transaction(ctx, updateFunc)
}

// transaction-wrapper-end
