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
	}

	fmt.Print("insert :d", status)

	resultJenisBarang.Status = "Success Insert data into Jenis Barang"

	return &resultJenisBarang, nil
}
