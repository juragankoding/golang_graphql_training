package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertOrders(ctx context.Context, custormerID int, orderStatus int, orderDate string, requiredDate string, shippedDate string, storeID int, staffID int) (*model.ResultInsertOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOrders(ctx context.Context, orderID int, custormerID int, orderStatus int, orderDate string, requiredDate string, shippedDate string, storeID int, staffID int) (*model.ResultUpdateOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrders(ctx context.Context, id int) (*model.ResultDeleteOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllOrders(ctx context.Context) (*model.ResultFetchOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SingelOrders(ctx context.Context, id *int) (*model.ResultGetOrders, error) {
	panic(fmt.Errorf("not implemented"))
}
