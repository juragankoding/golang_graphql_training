package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertCustomers(ctx context.Context, nama string) (*model.ResultInsertCustomers, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCustomers(ctx context.Context, id int, nama string) (*model.ResultUpdateCustomers, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCustomers(ctx context.Context, id int) (*model.ResultDeleteCustomers, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customers(ctx context.Context) (*model.ResultFetchCustomers, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customer(ctx context.Context, id *int) (*model.ResultGetCustomers, error) {
	panic(fmt.Errorf("not implemented"))
}
