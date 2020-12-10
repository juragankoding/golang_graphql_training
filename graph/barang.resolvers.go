package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	generated1 "github.com/juragankoding/golang_graphql_training/generated"
	"github.com/juragankoding/golang_graphql_training/graph/model"
	"github.com/juragankoding/golang_graphql_training/models"
)

func (r *mutationResolver) InsertBarang(ctx context.Context, id int, nama string, description string, jenisBarang int) (*model.ResultInsertBarang, error) {
	barang := models.Barang{
		Nama:        nama,
		Description: description,
		Jumlah:      0,
		JenisBarang: jenisBarang,
	}

	resultInsertBarang := model.ResultInsertBarang{
		Status: "nothing do anything",
		Code:   404,
		Data:   nil,
	}

	statusID, err := barang.Insert()

	if err != nil {
		log.Fatal(err)

		resultInsertBarang.Status = "kesalahan insert database"

		return &resultInsertBarang, nil
	}

	resultInsertBarang.Code = 200
	resultInsertBarang.Status = fmt.Sprintf("Success insert barang: %d", statusID)
	resultInsertBarang.Data = &barang

	return &resultInsertBarang, nil
}

func (r *queryResolver) GetBarang(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
