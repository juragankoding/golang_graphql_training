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

func (r *mutationResolver) InsertStocks(ctx context.Context, storeID int, productID int, quantity int) (*model.ResultInsertStocks, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	stoks := domain.Stocks{
		StoreID:   storeID,
		ProductID: productID,
		Quantity:  quantity,
	}

	_, err := r.StocksUseCase.Insert(stoks)

	if err != nil {
		return nil, err
	}

	return &model.ResultInsertStocks{
		Status: "success",
		Code:   200,
		Data:   &stoks,
	}, nil
}

func (r *mutationResolver) UpdateStocks(ctx context.Context, storeID int, productID int, quantity int) (*model.ResultUpdateStocks, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	stoks := domain.Stocks{
		StoreID:   storeID,
		ProductID: productID,
		Quantity:  quantity,
	}

	countEffect, err := r.StocksUseCase.Update(stoks)

	if err != nil {
		return nil, err
	}

	if countEffect <= 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultUpdateStocks{
		Status: "success",
		Code:   200,
		Data:   &stoks,
	}, nil
}

func (r *mutationResolver) DeleteStocks(ctx context.Context, storeID int, productID int) (*model.ResultDeleteStocks, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	countEffect, err := r.StocksUseCase.Delete(storeID, productID)

	if err != nil {
		return nil, err
	}

	if countEffect <= 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultDeleteStocks{
		Status: "success",
		Code:   200,
	}, nil
}

func (r *queryResolver) AllStocks(ctx context.Context) (*model.ResultFetchStocks, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	stocks, err := r.StocksUseCase.All()

	if err != nil {
		return nil, err
	}

	return &model.ResultFetchStocks{
		Status: "success",
		Code:   200,
		Data:   stocks,
	}, nil
}

func (r *queryResolver) SingelStocks(ctx context.Context, storeID int, productID int) (*model.ResultGetStocks, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	stocks, err := r.StocksUseCase.Single(storeID, productID)

	if err != nil {
		return nil, err
	}

	return &model.ResultGetStocks{
		Status: "success",
		Code:   200,
		Data:   stocks,
	}, nil
}

func (r *stocksResolver) ProducData(ctx context.Context, obj *domain.Stocks) (*domain.Products, error) {
	panic(fmt.Errorf("not implemented"))
}

// Stocks returns generated.StocksResolver implementation.
func (r *Resolver) Stocks() generated.StocksResolver { return &stocksResolver{r} }

type stocksResolver struct{ *Resolver }
