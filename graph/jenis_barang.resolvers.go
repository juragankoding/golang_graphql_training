package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/juragankoding/golang_graphql_training/model"
	"github.com/juragankoding/golang_graphql_training/models"
)

func (r *mutationResolver) InsertJenisBarang(ctx context.Context, jenisBarang string) (*model.ResultJenisBarang, error) {
	resultJenisBarang := model.ResultJenisBarang{}

	jb := models.JenisBarang{
		ID:          0,
		JenisBarang: jenisBarang,
	}

	status, error := jb.InsertJenisBarang()

	if error != nil {
		log.Fatal(error)

		resultJenisBarang.Code = 404
		resultJenisBarang.Status = "Insert to database not success for Jenis Barang"

		return &resultJenisBarang, error
	}

	fmt.Printf("insert %d", status)

	resultJenisBarang.Code = 200
	resultJenisBarang.Status = "Success Insert data into Jenis Barang"
	resultJenisBarang.Data = &models.JenisBarang{
		ID:          int(status),
		JenisBarang: jenisBarang,
	}

	return &resultJenisBarang, nil
}

func (r *queryResolver) GetAllJenisBarang(ctx context.Context) (*model.ResultGetAllJenisBarang, error) {
	resultGetAllJenisBarang := model.ResultGetAllJenisBarang{
		Status: "nothing happens",
		Code:   0,
		Data:   nil,
	}

	data, err := models.GetAll()

	if err != nil {
		log.Fatal(err)

		resultGetAllJenisBarang.Code = 404
		resultGetAllJenisBarang.Status = err.Error()

		return &resultGetAllJenisBarang, err
	}

	resultGetAllJenisBarang.Code = 200
	resultGetAllJenisBarang.Status = "Success gettering data"
	resultGetAllJenisBarang.Data = data

	return &resultGetAllJenisBarang, nil
}
