package main

import (
	"context"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// app-start
type PgTables struct {
	SchemaName string `gorm:"column:schemaname"`
	TableName  string `gorm:"column:tablename"`
}

func main() {
	// Gormのロガーのファクトリー関数を用いて生成
	newLogger := logger.New(
		log.New(os.Stdout, "\r", log.LstdFlags),
		logger.Config{LogLevel: logger.Info},
	)

	db, err := gorm.Open(postgres.Open("user=testuser password=pass host=localhost port=5432 dbname=testdb sslmode=disable"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	var pgtables []PgTables
	ctx := context.Background()
	if err := db.WithContext(ctx).Where("schemaname = ?", "information_schema").Find(&pgtables).Error; err != nil {
		log.Fatal(err)
	}
}

// app-end
