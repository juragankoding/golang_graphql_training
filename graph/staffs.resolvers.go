package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/domain"
	"github.com/juragankoding/golang_graphql_training/generated"
	"github.com/juragankoding/golang_graphql_training/model"
)

func (r *mutationResolver) InsertStaffs(ctx context.Context, firstName string, lastName string, phone string, email string, active int, storeID int, managerID int) (*model.ResultInsertStaffs, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStaffs(ctx context.Context, staffID int, firstName string, lastName string, phone string, email string, active int, storeID int, managerID int) (*model.ResultUpdateStaffs, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteStaffs(ctx context.Context, id int) (*model.ResultDeleteStaffs, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllStaffs(ctx context.Context) (*model.ResultFetchStaffs, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SingelStaffs(ctx context.Context, id *int) (*model.ResultGetStaffs, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *staffsResolver) Active(ctx context.Context, obj *domain.Staffs) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *staffsResolver) StoreData(ctx context.Context, obj *domain.Staffs) (*model.Stores, error) {
	panic(fmt.Errorf("not implemented"))
}

// Staffs returns generated.StaffsResolver implementation.
func (r *Resolver) Staffs() generated.StaffsResolver { return &staffsResolver{r} }

type staffsResolver struct{ *Resolver }
