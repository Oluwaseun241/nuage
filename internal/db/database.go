package db

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func InitDB() (*bun.DB, error) {
	sqlite, err := sql.Open(sqliteshim.ShimName, "data/sqlite.db")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqlite, sqlitedialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping database: %v", err))
	}
	return db, nil
}
