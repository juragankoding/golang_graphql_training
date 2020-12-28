package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/juragankoding/golang_graphql_training/domain"
	errorJurganKoding "github.com/juragankoding/golang_graphql_training/error"
	"github.com/juragankoding/golang_graphql_training/middleware"
	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertOrderItem(ctx context.Context, orderID int, productID int, quantity int, listPrice int, discount int) (*model.ResultInsertOrderItem, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	orderITEM := domain.OrderItem{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		ListPrice: listPrice,
		Discount:  discount,
	}

	lastInsertedID, err := r.OrderItemUseCase.Insert(orderITEM)

	if err != nil {
		return nil, err
	}

	orderITEM.ItemID = int(lastInsertedID)

	return &model.ResultInsertOrderItem{
		Status: "success",
		Code:   200,
		Data:   &orderITEM,
	}, nil
}

func (r *mutationResolver) UpdateOrderItem(ctx context.Context, itemID int, orderID int, productID int, quantity int, listPrice int, discount int) (*model.ResultUpdateOrderItem, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	orderITEM := domain.OrderItem{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		ListPrice: listPrice,
		Discount:  discount,
	}

	rowEffect, err := r.OrderItemUseCase.Update(orderITEM)

	if err != nil {
		return nil, err
	}

	if rowEffect == 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultUpdateOrderItem{
		Status: "success",
		Code:   200,
		Data:   &orderITEM,
	}, nil
}

func (r *mutationResolver) DeleteOrderItem(ctx context.Context, id int) (*model.ResultDeleteOrderItem, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	rowEffect, err := r.OrderItemUseCase.Delete(id)

	if err != nil {
		return nil, err
	}

	if rowEffect == 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultDeleteOrderItem{
		Status: "success",
		Code:   200,
	}, nil
}

func (r *queryResolver) AllOrderItem(ctx context.Context) (*model.ResultFetchOrderItem, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	customers, err := r.OrderItemUseCase.Fetch()

	if err != nil {
		return nil, err
	}

	return &model.ResultFetchOrderItem{
		Status: "success",
		Code:   200,
		Data:   customers,
	}, nil
}

func (r *queryResolver) SingelOrderItem(ctx context.Context, id *int) (*model.ResultGetOrderItem, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	customer, err := r.OrderItemUseCase.Get(*id)

	if err != nil {
		return nil, err
	}

	return &model.ResultGetOrderItem{
		Status: "success",
		Code:   200,
		Data:   customer,
	}, nil
}
