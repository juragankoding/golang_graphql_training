package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/graph/generated"
	"github.com/juragankoding/golang_graphql_training/models"
)

func (r *queryResolver) GetBarang(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) InsertBarang(ctx context.Context, id int, nama string, description string) (*models.Barang, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
