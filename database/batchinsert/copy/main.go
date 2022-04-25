package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v4"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// copy-start
	conn, err := pgx.Connect(context.Background(), "postgres://testuser:pass@localhost:5432/testdb")
	if err != nil {
		log.Fatal(err)
	}

	txn, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}

	rows := [][]interface{}{
		{1, "おにぎり", 120},
		{2, "パン", 300},
		{3, "お茶", 100},
	}

	_, err = txn.CopyFrom(
		ctx, pgx.Identifier{"products"},
		[]string{"product_no", "name", "price"}, pgx.CopyFromRows(rows),
	)
	// copy-end

	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
