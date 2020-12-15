package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/juragankoding/golang_graphql_training/domain"
	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) Categories(ctx context.Context, nama string) (*model.ResultInsertCategories, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) InsertCategories(ctx context.Context, nama string) (*model.ResultInsertCategories, error) {
	resultInsertCategories := model.ResultInsertCategories{
		Status: "",
		Code:   404,
		Data:   nil,
	}

	categories := domain.Categories{
		Name: nama,
	}

	lastId, _, err := r.DomainCategoriUseCase.Insert(categories)

	if err != nil {
		log.Fatal(err)

		return &resultInsertCategories, err
	}

	categories.ID = int(lastId)

	resultInsertCategories.Status = "Sucess insert categories"
	resultInsertCategories.Code = 200
	resultInsertCategories.Data = &categories

	return &resultInsertCategories, nil
}
func (r *queryResolver) Fetch(ctx context.Context) (*model.ResultFetchCategories, error) {
	panic(fmt.Errorf("not implemented"))
}
