package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertStores(ctx context.Context, storeName string, phone string, email string, city string, state string, zipCode string) (*model.ResultInsertStores, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStores(ctx context.Context, storeID int, storeName string, phone string, email string, city string, state string, zipCode string) (*model.ResultUpdateStores, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteStores(ctx context.Context, id int) (*model.ResultDeleteStores, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllStores(ctx context.Context) (*model.ResultAllStores, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SingleStores(ctx context.Context, id *int) (*model.ResultSingleStores, error) {
	panic(fmt.Errorf("not implemented"))
}
