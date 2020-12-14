package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertBrands(ctx context.Context, nama string) (*model.ResultInsertBrands, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBrands(ctx context.Context, id int, nama string) (*model.ResultUpdateBrands, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBrands(ctx context.Context, id int) (*model.ResultDeleteBrands, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SingleBrands(ctx context.Context, id int) (*model.ResultSingleBrands, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllBrands(ctx context.Context) (*model.ResultAllBrands, error) {
	panic(fmt.Errorf("not implemented"))
}
