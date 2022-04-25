package main

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type PgTables struct {
	SchemaName string `db:"schemaname"`
	TableName  string `db:"tablename"`
}

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	ctx := context.Background()

	db, err := sqlx.Connect("postgres", "user=testuser password=pass host=localhost port=5432 dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	// app-log-start
	var pgtables []PgTables
	sql := `SELECT schemaname, tablename FROM pg_tables WHERE schemaname = $1;`
	args := `information_schema`

	log.Info().Msgf("sql=[%s], args=[%v]", sql, args)
	if err := db.SelectContext(ctx, &pgtables, sql, args); err != nil {
		log.Fatal().Msgf("%v", err)
	}
	// app-log-end

}
