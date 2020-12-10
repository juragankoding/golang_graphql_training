package models

import (
	"fmt"
	"log"

	"github.com/juragankoding/golang_graphql_training/db"
)

type JenisBarang struct {
	ID          int    `json:"id"`
	JenisBarang string `json:"jenis_barnag"`
}

func (jenis *JenisBarang) InsertJenisBarang() (int64, error) {
	sql, err := db.GetDatabase()

	if err != nil {
		fmt.Printf("kesalahan di")

		return -1, err
	}

	query := fmt.Sprintf("INSERT INTO jenis_barang (jenis) values ('%s')", jenis.JenisBarang)
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

func GetAll() ([]*JenisBarang, error) {
	var jenisBarangs []*JenisBarang

	database, err := db.GetDatabase()

	if err != nil {
		fmt.Printf("Get database failed")

		log.Fatal(err)

		return nil, err
	}

	rows, err := database.Query("SELECT * FROM jenis_barang")

	if err != nil {
		fmt.Printf("Kesalahan di pengambilan data dari database")

		log.Fatal(err)

		return nil, err
	}

	var jenisBarangTmp JenisBarang

	for rows.Next() {
		rows.Scan(&jenisBarangTmp.ID, &jenisBarangTmp.JenisBarang)

		fmt.Print("Get data from rows with id ", jenisBarangTmp.ID, ", jenis barang : ", jenisBarangTmp.JenisBarang)

		jenisBarangs = append(jenisBarangs, &jenisBarangTmp)
	}

	return jenisBarangs, nil
}
