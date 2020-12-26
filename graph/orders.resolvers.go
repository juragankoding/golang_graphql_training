package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/domain"
	errorJurganKoding "github.com/juragankoding/golang_graphql_training/error"
	"github.com/juragankoding/golang_graphql_training/generated"
	"github.com/juragankoding/golang_graphql_training/middleware"
	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertOrders(ctx context.Context, custormerID int, orderStatus int, orderDate string, requiredDate string, shippedDate string, storeID int, staffID int) (*model.ResultInsertOrders, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	orders := domain.Orders{
		CustomerID:   custormerID,
		OrderStatus:  orderStatus,
		OrderDate:    orderDate,
		RequiredDate: requiredDate,
		ShippedDate:  shippedDate,
		StoreID:      storeID,
		StaffID:      staffID,
	}

	lastInsertID, err := r.OrdersUseCase.Insert(orders)

	if err != nil {
		return nil, err
	}

	orders.CustomerID = int(lastInsertID)

	return &model.ResultInsertOrders{
		Status: "success",
		Code:   200,
		Data:   &orders,
	}, nil
}

func (r *mutationResolver) UpdateOrders(ctx context.Context, orderID int, custormerID int, orderStatus int, orderDate string, requiredDate string, shippedDate string, storeID int, staffID int) (*model.ResultUpdateOrders, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	orders := domain.Orders{
		OrderID:      orderID,
		CustomerID:   custormerID,
		OrderStatus:  orderStatus,
		OrderDate:    orderDate,
		RequiredDate: requiredDate,
		ShippedDate:  shippedDate,
		StoreID:      storeID,
		StaffID:      staffID,
	}

	countEffectRow, err := r.OrdersUseCase.Update(orders)

	if err != nil {
		return nil, err
	}

	if countEffectRow <= 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultUpdateOrders{
		Status: "success",
		Code:   200,
		Data:   &orders,
	}, nil
}

func (r *mutationResolver) DeleteOrders(ctx context.Context, id int) (*model.ResultDeleteOrders, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	countEffectRow, err := r.OrdersUseCase.Delete(id)

	if err != nil {
		return nil, err
	}

	if countEffectRow <= 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultDeleteOrders{
		Status: "success",
		Code:   200,
	}, nil
}

func (r *ordersResolver) StoreData(ctx context.Context, obj *domain.Orders) (*domain.Stores, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *ordersResolver) StaffData(ctx context.Context, obj *domain.Orders) (*domain.Staffs, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllOrders(ctx context.Context) (*model.ResultFetchOrders, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	allOrders, err := r.OrdersUseCase.All()

	if err != nil {
		return nil, err
	}

	return &model.ResultFetchOrders{
		Status: "success",
		Code:   200,
		Data:   allOrders,
	}, nil
}

func (r *queryResolver) SingelOrders(ctx context.Context, id *int) (*model.ResultGetOrders, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	orders, err := r.OrdersUseCase.Single(*id)

	if err != nil {
		return nil, err
	}

	return &model.ResultGetOrders{
		Status: "success",
		Code:   200,
		Data:   orders,
	}, nil
}

// Orders returns generated.OrdersResolver implementation.
func (r *Resolver) Orders() generated.OrdersResolver { return &ordersResolver{r} }

type ordersResolver struct{ *Resolver }
