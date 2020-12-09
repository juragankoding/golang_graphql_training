package models

import (
	"fmt"

	"github.com/juragankoding/golang_graphql_training/db"
)

type Barang struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	Description string `json:"description"`
}

func insertBarangDatabase(barang Barang) {
	sql, err := db.GetDatabase()

	if err != nil {
		fmt.Printf("kesalahan di")
	}

	sql.Prepare(fmt.Sprint("INSERT INTO jenis_barang (jenis) values (:s)", barang.Nama))
}
