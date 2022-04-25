package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PgTables struct {
	SchemaName string `gorm:"column:schemaname"`
	TableName  string `gorm:"column:tablename"`
}

type dbLoggerFunc func(string, ...interface{})

func (f dbLoggerFunc) Printf(str string, v ...interface{}) {
	f(str, v)
}

func main() {
	zerolog.TimeFieldFormat = time.RFC3339

	// database-custom-log-start
	var dbLogger dbLoggerFunc = func(str string, v ...interface{}) {
		// zerologを使ってGormのログを出力できるようにする
		log.Info().Msgf(fmt.Sprintf("%v", v...))
	}
	newLogger := logger.New(
		dbLogger,
		logger.Config{LogLevel: logger.Info},
	)
	// database-custom-log-end

	db, err := gorm.Open(postgres.Open("user=testuser password=pass host=localhost port=5432 dbname=testdb sslmode=disable"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	var pgtables []PgTables
	ctx := context.Background()
	if err := db.WithContext(ctx).Where("schemaname = ?", "information_schema").Find(&pgtables).Error; err != nil {
		log.Fatal().Msgf("%v", err)
	}
	for _, v := range pgtables {
		log.Info().Msgf("%v", v)
	}
}
