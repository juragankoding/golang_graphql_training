package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/juragankoding/golang_graphql_training/domain"
	errorJurganKoding "github.com/juragankoding/golang_graphql_training/error"
	"github.com/juragankoding/golang_graphql_training/generated"
	"github.com/juragankoding/golang_graphql_training/middleware"
	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertBrands(ctx context.Context, nama string) (*model.ResultInsertBrands, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	brands := domain.Brands{
		Name: nama,
	}

	ID, err := r.BrandsUseCase.Insert(brands)

	brands.ID = int(ID)

	if err != nil {
		return nil, err
	}

	return &model.ResultInsertBrands{
		Status: "success",
		Code:   200,
		Data:   &brands,
	}, nil
}

func (r *mutationResolver) UpdateBrands(ctx context.Context, id int, nama string) (*model.ResultUpdateBrands, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	brands := domain.Brands{
		ID:   id,
		Name: nama,
	}

	statusCode, err := r.BrandsUseCase.Update(brands)

	if err != nil {
		return nil, err
	}

	if statusCode == 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultUpdateBrands{
		Status: "success",
		Code:   200,
		Data:   &brands,
	}, nil
}

func (r *mutationResolver) DeleteBrands(ctx context.Context, id int) (*model.ResultDeleteBrands, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	rowEffectCount, err := r.BrandsUseCase.Delete(id)

	if err != nil {
		return nil, err
	}

	if rowEffectCount > 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultDeleteBrands{
		Status: "success",
		Code:   200,
	}, nil
}

func (r *queryResolver) SingleBrands(ctx context.Context, id int) (*model.ResultSingleBrands, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	userSingle, err := r.BrandsUseCase.Single(id)

	if err != nil {
		return nil, err
	}

	return &model.ResultSingleBrands{
		Status: "success",
		Code:   200,
		Data:   userSingle,
	}, nil
}

func (r *queryResolver) AllBrands(ctx context.Context) (*model.ResultAllBrands, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	brands, err := r.BrandsUseCase.All()

	if err != nil {
		return nil, err
	}

	return &model.ResultAllBrands{
		Status: "success",
		Code:   200,
		Data:   brands,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
