package db

import (
	"database/sql"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/utils"
	_ "github.com/mattn/go-sqlite3"
)

func GetDatabase() (*sql.DB, error) {
	fmt.Print(utils.RootDir(), "report.db")

	return sql.Open("sqlite3", fmt.Sprint(utils.RootDir(), "/report.db"))
}
