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

func (r *mutationResolver) InsertProducts(ctx context.Context, name string, brandID int, categoryID int, modelyear int, listPrice int) (*model.ResultInsertProducts, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProducts(ctx context.Context, id int, name string, brandID int, categoryID int, modelyear int, listPrice int) (*model.ResultUpdateProducts, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProdcuts(ctx context.Context, id int) (*model.ResultDeleteProducts, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productsResolver) ProductName(ctx context.Context, obj *domain.Products) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllProducts(ctx context.Context) (*model.ResultAllProducts, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SingleProducts(ctx context.Context, id *int) (*model.ResultSingleProducts, error) {
	panic(fmt.Errorf("not implemented"))
}

// Products returns generated.ProductsResolver implementation.
func (r *Resolver) Products() generated.ProductsResolver { return &productsResolver{r} }

type productsResolver struct{ *Resolver }
