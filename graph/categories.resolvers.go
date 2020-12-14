package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/juragankoding/golang_graphql_training/domain"
	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertCategories(ctx context.Context, nama string) (*model.ResultInsertCategories, error) {
	resultInsertCategories := model.ResultInsertCategories{
		Status: "",
		Code:   404,
		Data:   nil,
	}

	categories := domain.Categories{
		Name: nama,
	}

	lastId, _, err := r.DomainCategoriUseCase.Insert(categories)

	if err != nil {
		log.Fatal(err)

		return &resultInsertCategories, err
	}

	categories.ID = int(lastId)

	resultInsertCategories.Status = "Sucess insert categories"
	resultInsertCategories.Code = 200
	resultInsertCategories.Data = &categories

	return &resultInsertCategories, nil
}

func (r *mutationResolver) UpdateCategories(ctx context.Context, id int, nama string) (*model.ResultUpdateCategories, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCategories(ctx context.Context, id int) (*model.ResultDeleteCategories, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Fetch(ctx context.Context) (*model.ResultFetchCategories, error) {
	panic(fmt.Errorf("not implemented"))
}
