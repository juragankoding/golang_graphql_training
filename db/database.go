package db

import (
	"database/sql"
)

func GetDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", "./report.db")
}
