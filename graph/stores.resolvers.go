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

func (r *mutationResolver) InsertStores(ctx context.Context, storeName string, phone string, email string, city string, state string, zipCode string) (*model.ResultInsertStores, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	stores := domain.Stores{
		StoreName: storeName,
		Phone:     phone,
		Email:     email,
		City:      city,
		State:     state,
		ZipCode:   zipCode,
	}

	lastInsertID, err := r.StoreUseCase.Insert(stores)

	if err != nil {
		return nil, err
	}

	stores.StoreID = int(lastInsertID)

	return &model.ResultInsertStores{
		Status: "success",
		Code:   200,
		Data:   &stores,
	}, nil
}

func (r *mutationResolver) UpdateStores(ctx context.Context, storeID int, storeName string, phone string, email string, city string, state string, zipCode string) (*model.ResultUpdateStores, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	stores := domain.Stores{
		StoreID:   storeID,
		StoreName: storeName,
		Phone:     phone,
		Email:     email,
		City:      city,
		State:     state,
		ZipCode:   zipCode,
	}

	countEffect, err := r.StoreUseCase.Update(stores)

	if err != nil {
		return nil, err
	}

	if countEffect <= 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultUpdateStores{
		Status: "success",
		Code:   200,
		Data:   &stores,
	}, nil
}

func (r *mutationResolver) DeleteStores(ctx context.Context, id int) (*model.ResultDeleteStores, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	countEffect, err := r.StoreUseCase.Delete(id)

	if err != nil {
		return nil, err
	}

	if countEffect <= 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultDeleteStores{
		Status: "success",
		Code:   200,
	}, nil
}

func (r *queryResolver) AllStores(ctx context.Context) (*model.ResultAllStores, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	stores, err := r.StoreUseCase.Fetch()

	if err != nil {
		return nil, err
	}

	return &model.ResultAllStores{
		Status: "success",
		Code:   200,
		Data:   stores,
	}, nil
}

func (r *queryResolver) SingleStores(ctx context.Context, id *int) (*model.ResultSingleStores, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	stores, err := r.StoreUseCase.Get(*id)

	if err != nil {
		return nil, err
	}

	return &model.ResultSingleStores{
		Status: "success",
		Code:   200,
		Data:   stores,
	}, nil
}
