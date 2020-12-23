package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.ResultLogin, error) {
	panic(fmt.Errorf("not implemented"))

}

func (r *mutationResolver) CreateUser(ctx context.Context, username string, password string) (*model.ResultCreateUser, error) {
	user, _, err := r.UserUseCase.CreateUser(&username, &password, &username)

	if err != nil {
		return nil, err
	}

	obj := model.ResultCreateUser{
		Data: "successs user",
		Code: 200,
		User: user,
	}

	return &obj, nil

}

func (r *queryResolver) ListUsers(ctx context.Context) (*model.ResultUsers, error) {
	users, err := r.UserUseCase.ListUsers()

	if err != nil {
		return nil, err
	}

	return &model.ResultUsers{
		Data: "",
		User: users,
		Code: 200,
	}, nil
}
