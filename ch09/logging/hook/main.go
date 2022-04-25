package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/gchaincl/sqlhooks"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type PgTable struct {
	SchemaName string `db:"schemaname"`
	TableName  string `db:"tablename"`
}

// hooks-start
var _ sqlhooks.Hooks = (*hook)(nil)

type hook struct{}

func (h *hook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	log.Printf("SQL:\n%v\nARGS:%v\n", query, args)
	return ctx, nil
}

func (h *hook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return ctx, nil
}

// hooks-end

// hook-app-start
func main() {
	ctx := context.Background()

	// "postgres-proxy" としてフックできるドライバーを登録
	sql.Register("postgres-proxy", sqlhooks.Wrap(stdlib.GetDefaultDriver(), &hook{}))
	db, err := sqlx.Connect("postgres-proxy", "user=testuser password=pass host=localhost port=5432 dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatalf("connect: %v\n", err)
	}

	var pgtables []PgTable
	sql := `SELECT schemaname, tablename FROM pg_tables WHERE schemaname = $1;`
	args := `information_schema`

	if err := db.SelectContext(ctx, &pgtables, sql, args); err != nil {
		log.Fatalf("select: %v\n", err)
	}
}

// hook-app-end
