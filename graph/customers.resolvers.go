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

func (r *mutationResolver) InsertCustomers(ctx context.Context, firstName string, lastName string, phone string, email string, street string, city string, state string, zipCode int) (*model.ResultInsertCustomers, error) {
	user := middleware.GetUserFromContext(ctx)

	if user != nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	customers := domain.Customers{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		Street:    street,
		City:      city,
		State:     state,
		ZipCode:   zipCode,
	}

	lastInsertID, err := r.CustomersUseCase.Insert(customers)

	if err != nil {
		return nil, err
	}

	customers.CustomerID = int(lastInsertID)

	return &model.ResultInsertCustomers{
		Status: "Success",
		Code:   200,
		Data:   &customers,
	}, nil
}

func (r *mutationResolver) UpdateCustomers(ctx context.Context, customerID int, firstName string, lastName string, phone string, email string, street string, city string, state string, zipCode int) (*model.ResultUpdateCustomers, error) {
	user := middleware.GetUserFromContext(ctx)

	if user != nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	customers := domain.Customers{
		FirstName:  firstName,
		LastName:   lastName,
		Phone:      phone,
		Email:      email,
		Street:     street,
		City:       city,
		State:      state,
		ZipCode:    zipCode,
		CustomerID: customerID,
	}

	countRowEffect, err := r.CustomersUseCase.Update(customers)

	if err != nil {
		return nil, err
	}

	if countRowEffect > 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultUpdateCustomers{
		Status: "Success",
		Code:   200,
		Data:   &customers,
	}, nil
}

func (r *mutationResolver) DeleteCustomers(ctx context.Context, id int) (*model.ResultDeleteCustomers, error) {
	user := middleware.GetUserFromContext(ctx)

	if user != nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	countRowEffect, err := r.CustomersUseCase.Delete(id)

	if err != nil {
		return nil, err
	}

	if countRowEffect > 0 {
		return nil, errorJurganKoding.ErrorsUpdateNotEffect
	}

	return &model.ResultDeleteCustomers{
		Status: "Success",
		Code:   200,
	}, nil
}

func (r *queryResolver) Customers(ctx context.Context) (*model.ResultFetchCustomers, error) {
	user := middleware.GetUserFromContext(ctx)

	if user != nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	customers, err := r.CustomersUseCase.All()

	if err != nil {
		return nil, err
	}

	return &model.ResultFetchCustomers{
		Status: "success",
		Code:   200,
		Data:   customers,
	}, nil
}

func (r *queryResolver) Customer(ctx context.Context, id *int) (*model.ResultGetCustomers, error) {
	user := middleware.GetUserFromContext(ctx)

	if user != nil {
		return nil, errorJurganKoding.ErrorsNotAuthenticate
	}

	customers, err := r.CustomersUseCase.Get(*id)

	if err != nil {
		return nil, err
	}

	return &model.ResultGetCustomers{
		Status: "success",
		Code:   200,
		Data:   customers,
	}, nil
}
