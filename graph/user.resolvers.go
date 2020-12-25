package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/juragankoding/golang_graphql_training/domain"
	"github.com/juragankoding/golang_graphql_training/middleware"
	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.ResultLogin, error) {
	_, token, err := r.UserUseCase.Login(&username, &password)

	if err != nil {
		return nil, err
	}

	return &model.ResultLogin{
		Token: token,
		Data:  "success login",
		Code:  200,
		Type:  "",
	}, nil
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

func (r *queryResolver) Users(ctx context.Context) (*domain.User, error) {
	userAuth := middleware.GetUserFromContext(ctx)

	user, err := r.UserUseCase.SingleUserFromID(userAuth.UserID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
