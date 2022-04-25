package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/gchaincl/sqlhooks"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type PgTables struct {
	SchemaName string `db:"schemaname"`
	TableName  string `db:"tablename"`
}

type hook struct{}

func (h *hook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	zlog.Info().Msgf("sql=[%s], args=%v", query, args)
	return ctx, nil
}

func (h *hook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return ctx, nil
}

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	ctx := context.Background()

	sql.Register("postgres-proxy", sqlhooks.Wrap(&pq.Driver{}, &hook{}))
	db, err := sqlx.Connect("postgres-proxy", "user=testuser password=pass host=localhost port=5432 dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}

	var pgtables []PgTables
	sql := `SELECT schemaname, tablename FROM pg_tables WHERE schemaname = $1;`
	args := `information_schema`

	if err := db.SelectContext(ctx, &pgtables, sql, args); err != nil {
		panic(err)
	}
}
