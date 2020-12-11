package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertCategories(ctx context.Context, nama string) (*model.ResultInsertCategories, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCategories(ctx context.Context, id int, nama string) (*model.ResultUpdateCategories, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCategories(ctx context.Context, id int) (*model.ResultDeleteCategories, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Fetch(ctx context.Context) (*model.ResultFetchCategories, error) {
	data, message, err := r.DomainCategoriUseCase.Fetch()

	panic(fmt.Errorf("not implemented"))
}
