package db

import (
	"context"
	"database/sql"
	"fmt"
	"nuage/migrations"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"
)

type SQLiteUserRepo struct {
	db *bun.DB
}

func NewSQLiteUserRepo() (*SQLiteUserRepo, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	return &SQLiteUserRepo{db: db}, nil
}

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

	if err := runMigration(db); err != nil {
		panic(fmt.Sprintf("failed to run migrations: %v", err))
	}

	return db, nil
}

func runMigration(db *bun.DB) error {
	ctx := context.Background()
	migrator := migrate.NewMigrator(db, migrations.Migrations)
	if err := migrator.Init(ctx); err != nil {
		return err
	}

	if err := migrator.Lock(ctx); err != nil {
		return err
	}
	defer migrator.Unlock(ctx)

	groups, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	if groups.IsZero() {
		fmt.Println("no new migrations to run")
		return nil
	}

	fmt.Printf("migrated to %s\n", groups.String())
	return nil
}
