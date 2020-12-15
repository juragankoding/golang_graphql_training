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

func (r *mutationResolver) InsertStocks(ctx context.Context, productID int, quantity int) (*model.ResultInsertStocks, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStocks(ctx context.Context, stockID int, productID int, quantity int) (*model.ResultUpdateStocks, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteStocks(ctx context.Context, id int) (*model.ResultDeleteStocks, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllStocks(ctx context.Context) (*model.ResultFetchStocks, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SingelStocks(ctx context.Context, id *int) (*model.ResultGetStocks, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stocksResolver) ProducData(ctx context.Context, obj *domain.Stocks) (*domain.Products, error) {
	panic(fmt.Errorf("not implemented"))
}

// Stocks returns generated.StocksResolver implementation.
func (r *Resolver) Stocks() generated.StocksResolver { return &stocksResolver{r} }

type stocksResolver struct{ *Resolver }
