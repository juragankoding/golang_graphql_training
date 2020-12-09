package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/juragankoding/golang_graphql_training/graph/model"
	"github.com/juragankoding/golang_graphql_training/models"
)

func (r *queryResolver) InsertJenisBarang(ctx context.Context, jenisBarang string) (*model.ResultJenisBarang, error) {
	resultJenisBarang := model.ResultJenisBarang{}
	status, error := models.InsertJenisBarang(
		models.JenisBarang{
			JenisBarang: jenisBarang,
			ID:          0,
		},
	)

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
	panic(fmt.Errorf("not implemented"))
}
