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
	resultInserCategories := model.ResultInsertCategories{
		Status: "",
		Code:   404,
		Data:   nil,
	}

	categories := domain.Categories{
		Name: nama,
	}

	lastID, _, err := r.DomainCategoriUseCase.Insert(categories)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	categories.ID = int(lastID)

	resultInserCategories.Data = &categories
	resultInserCategories.Code = 200
	resultInserCategories.Status = "Success insert into categories"

	return &resultInserCategories, nil
}

func (r *mutationResolver) UpdateCategories(ctx context.Context, id int, nama string) (*model.ResultUpdateCategories, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCategories(ctx context.Context, id int) (*model.ResultDeleteCategories, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllCategories(ctx context.Context) (*model.ResultGetAllCategories, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCategories(ctx context.Context, id int) (*model.ResultGetCategories, error) {
	panic(fmt.Errorf("not implemented"))
}
