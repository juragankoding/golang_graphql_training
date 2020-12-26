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

func (r *mutationResolver) InsertStaffs(ctx context.Context, firstName string, lastName string, phone string, email string, active string, storeID int, managerID int) (*model.ResultInsertStaffs, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	staffs := domain.Staffs{
		StoreID:   storeID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Active:    active,
		ManagerID: managerID,
	}

	lastInsertID, err := r.StaffsUseCase.Insert(staffs)

	if err != nil {
		return nil, err
	}

	staffs.StaffID = int(lastInsertID)

	return &model.ResultInsertStaffs{
		Status: "success",
		Code:   200,
		Data:   &staffs,
	}, nil
}

func (r *mutationResolver) UpdateStaffs(ctx context.Context, staffID int, firstName string, lastName string, phone string, email string, active string, storeID int, managerID int) (*model.ResultUpdateStaffs, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	staffs := domain.Staffs{
		StoreID:   storeID,
		StaffID:   staffID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Active:    active,
		ManagerID: managerID,
	}

	rwoEffectCount, err := r.StaffsUseCase.Update(staffs)

	if err != nil {
		return nil, err
	}

	if rwoEffectCount <= 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultUpdateStaffs{
		Status: "success",
		Code:   200,
		Data:   &staffs,
	}, nil
}

func (r *mutationResolver) DeleteStaffs(ctx context.Context, id int) (*model.ResultDeleteStaffs, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	rwoEffectCount, err := r.StaffsUseCase.Delete(id)

	if err != nil {
		return nil, err
	}

	if rwoEffectCount <= 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultDeleteStaffs{
		Status: "success",
		Code:   200,
	}, nil
}

func (r *queryResolver) AllStaffs(ctx context.Context) (*model.ResultFetchStaffs, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	customers, err := r.StaffsUseCase.All()

	if err != nil {
		return nil, err
	}

	return &model.ResultFetchStaffs{
		Status: "success",
		Code:   200,
		Data:   customers,
	}, nil
}

func (r *queryResolver) SingelStaffs(ctx context.Context, id *int) (*model.ResultGetStaffs, error) {
	user := middleware.GetUserFromContext(ctx)

	if user == nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	customers, err := r.StaffsUseCase.Single(*id)

	if err != nil {
		return nil, err
	}

	return &model.ResultGetStaffs{
		Status: "success",
		Code:   200,
		Data:   customers,
	}, nil
}
