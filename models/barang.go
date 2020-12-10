package models

import (
	"fmt"
	"log"

	"github.com/juragankoding/golang_graphql_training/db"
)

type Barang struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	Description string `json:"description"`
	Jumlah      int    `json:"jumlah"`
	JenisBarang int    `json:"jenisBarang"`
}

func (barang *Barang) Insert() (int64, error) {
	database, error := db.GetDatabase()

	if error != nil {
		log.Fatal(error)

		return -1, error
	}

	stmt, err := database.Prepare("INSERT INTO barang (nama, desc, jumlah, jenis_barang) VALUES (?,?,?,?)")

	if err != nil {
		log.Fatal(err)

		return -1, err
	}

	statementStatus, err := stmt.Exec(barang.Nama, barang.Description, barang.Jumlah, barang.JenisBarang)

	if err != nil {
		fmt.Printf("Keslahaan di ekseskusi Statement")

		return -1, err
	}

	id, err := statementStatus.LastInsertId()

	if err != nil {
		log.Fatal(err)

		return -1, err
	}

	barang.ID = int(id)

	return (id), nil
}
