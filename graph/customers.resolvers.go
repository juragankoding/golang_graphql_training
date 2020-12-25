package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

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

	lastInsertId, err := r.CustomersUseCase.Insert(customers)

	if err != nil {
		return nil, err
	}

	customers.CustomerID = int(lastInsertId)

	return &model.ResultInsertCustomers{
		Status: "Success",
		Code:   200,
		Data:   &customers,
	}, nil
}

func (r *mutationResolver) UpdateCustomers(ctx context.Context, customerID int, firstName string, lastName string, phone string, email string, street string, city string, state string, zipCode int) (*model.ResultUpdateCustomers, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCustomers(ctx context.Context, id int) (*model.ResultDeleteCustomers, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customers(ctx context.Context) (*model.ResultFetchCustomers, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customer(ctx context.Context, id *int) (*model.ResultGetCustomers, error) {
	panic(fmt.Errorf("not implemented"))
}
