package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", "./report.db")
}
