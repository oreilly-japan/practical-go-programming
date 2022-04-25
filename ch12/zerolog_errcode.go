package main

import (
	"database/sql"
	"github.com/rs/zerolog/log"
)

// errcode-start
const DBConnErrCD = "E10001"

func businessLogic() {
	db, err := openDB()
	if err != nil {
		// エラーコード付きで出力
		log.Error().Str("code", DBConnErrCD).Err(err).Msg("db connection failed")
	}
	// ...
}

// errcode-end

func openDB() (*sql.DB, error) {
	return nil, nil
}
