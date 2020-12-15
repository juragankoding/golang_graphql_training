package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/juragankoding/golang_graphql_training/domain"
	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) Categories(ctx context.Context, nama string) (*model.ResultInsertCategories, error) {
	resultInserCategories := model.ResultInsertCategories{
		Status: "",
		Code:   404,
		Data:   nil,
	}

	categories := domain.Categories{
		Name: nama,
	}

	lastID, _, err := r.DomainCategoriUseCase.Insert(categories)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	categories.ID = int(lastID)

	resultInserCategories.Data = &categories
	resultInserCategories.Code = 200
	resultInserCategories.Status = "Success insert into categories"

	return &resultInserCategories, nil
}

func (r *mutationResolver) UpdateCategories(ctx context.Context, id int, nama string) (*model.ResultUpdateCategories, error) {
	resultUpdateCategoreis := model.ResultUpdateCategories{
		Status: "",
		Code:   404,
		Data:   nil,
	}

	categories := domain.Categories{
		ID:   id,
		Name: nama,
	}

	_, err := r.DomainCategoriUseCase.Update(categories)

	if err != nil {
		return nil, err
	}

	resultUpdateCategoreis.Code = 200
	resultUpdateCategoreis.Status = "Success update categories"
	resultUpdateCategoreis.Data = &categories

	return &resultUpdateCategoreis, nil
}

func (r *mutationResolver) DeleteCategories(ctx context.Context, id int) (*model.ResultDeleteCategories, error) {
	_, err := r.DomainCategoriUseCase.Delete(id)

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	return &model.ResultDeleteCategories{
		Status: "Success Delete Categories",
		Code:   200,
	}, nil
}

func (r *queryResolver) Categories(ctx context.Context) (*model.ResultFetchCategories, error) {
	var categories []*domain.Categories
	var err error

	categories, _, err = r.DomainCategoriUseCase.Fetch()

	if err != nil {
		log.Fatal(err)
	}

	return &model.ResultFetchCategories{
		Status: "Success Gettering Data",
		Code:   200,
		Data:   categories,
	}, nil
}

func (r *queryResolver) Category(ctx context.Context, id *int) (*model.ResultGetCategories, error) {
	var categories *domain.Categories
	var err error

	categories, _, err = r.DomainCategoriUseCase.Get(*id)

	if err != nil {
		return nil, err
	}

	return &model.ResultGetCategories{
		Status: "Success Gettering data",
		Code:   200,
		Data:   categories,
	}, nil
}
