package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/juragankoding/golang_graphql_training/domain"
	errorJurganKoding "github.com/juragankoding/golang_graphql_training/error"
	"github.com/juragankoding/golang_graphql_training/middleware"
	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertProducts(ctx context.Context, name string, brandID int, categoryID int, modelyear int, listPrice int) (*model.ResultInsertProducts, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	products := domain.Products{
		ProductName: name,
		BrandID:     brandID,
		CategoryID:  categoryID,
		ModelYear:   strconv.Itoa(modelyear),
		ListPrice:   listPrice,
	}
	lastInsertID, err := r.ProductUseCase.Insert(products)

	if err != nil {
		return nil, err
	}

	products.ProductID = int(lastInsertID)

	return &model.ResultInsertProducts{
		Status: "success",
		Code:   200,
		Data:   &products,
	}, nil
}

func (r *mutationResolver) UpdateProducts(ctx context.Context, id int, name string, brandID int, categoryID int, modelyear int, listPrice int) (*model.ResultUpdateProducts, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	products := domain.Products{
		ProductName: name,
		BrandID:     brandID,
		CategoryID:  categoryID,
		ModelYear:   strconv.Itoa(modelyear),
		ListPrice:   listPrice,
		ProductID:   id,
	}

	rowEffectID, err := r.ProductUseCase.Insert(products)

	if err != nil {
		return nil, err
	}

	if rowEffectID == 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultUpdateProducts{
		Status: "success",
		Code:   200,
		Data:   &products,
	}, nil
}

func (r *mutationResolver) DeleteProducts(ctx context.Context, id int) (*model.ResultDeleteProducts, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	rowEffectID, err := r.ProductUseCase.Delete(id)

	if err != nil {
		return nil, err
	}

	if rowEffectID == 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultDeleteProducts{
		Status: "success",
		Code:   200,
	}, nil
}

func (r *queryResolver) AllProducts(ctx context.Context) (*model.ResultAllProducts, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	products, err := r.ProductUseCase.All()

	if err != nil {
		return nil, err
	}

	return &model.ResultAllProducts{
		Status: "success",
		Code:   200,
		Data:   products,
	}, nil
}

func (r *queryResolver) SingleProducts(ctx context.Context, id *int) (*model.ResultSingleProducts, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	products, err := r.ProductUseCase.Single(*id)

	if err != nil {
		return nil, err
	}

	return &model.ResultSingleProducts{
		Status: "success",
		Code:   200,
		Data:   products,
	}, nil
}
