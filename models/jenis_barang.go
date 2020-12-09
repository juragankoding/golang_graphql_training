package models

import (
	"fmt"

	"github.com/juragankoding/golang_graphql_training/db"
)

type JenisBarang struct {
	ID          int    `json:"id"`
	JenisBarang string `json:"jenis_barnag"`
}

func InsertJenisBarang(barang JenisBarang) (int64, error) {
	sql, err := db.GetDatabase()

	if err != nil {
		fmt.Printf("kesalahan di")

		return -1, err
	}

	query := fmt.Sprintf("INSERT INTO jenis_barang (jenis) values ('%s')", barang.JenisBarang)
	fmt.Print(query)
	statement, err := sql.Prepare(query)

	if err != nil {
		fmt.Printf("Kesalahan di builder statement")

		return -1, err
	}

	res, err := statement.Exec()

	if err != nil {
		fmt.Printf("kesalahan di ekseskusi statement")

		return -1, err
	}
	return res.LastInsertId()
}

func GetJenisBarangSingle() ([]JenisBarang, error) {
	var jenisBarang []JenisBarang

	return jenisBarang, nil
}
