package db

import (
	"database/sql"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/utils"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprint(utils.RootDir(), "/report.db"))

	if err != nil {
		return db, err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: fmt.Sprint(utils.RootDir(), "/migrations/sqlite3"),
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Applied %d migrations!\n", n)

	return db, nil
}
