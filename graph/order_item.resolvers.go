package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertOrderItem(ctx context.Context, orderID int, productID int, quantity int, listPrice int, discount int) (*model.ResultInsertOrderItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOrderItem(ctx context.Context, itemID int, orderID int, productID int, quantity int, listPrice int, discount int) (*model.ResultUpdateOrderItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrderItem(ctx context.Context, id int) (*model.ResultDeleteOrderItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllOrderItem(ctx context.Context) (*model.ResultFetchOrderItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SingelOrderItem(ctx context.Context, id *int) (*model.ResultGetOrderItem, error) {
	panic(fmt.Errorf("not implemented"))
}
