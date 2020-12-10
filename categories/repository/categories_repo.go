package repository

import (
	"context"
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type categoriesRepository struct {
	Conn *sql.DB
}

func NewCategoriesRepository(Conn *sql.DB) domain.CategoriesRepository {
	return &categoriesRepository{
		Conn: Conn,
	}
}

func (r *categoriesRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Categories, nextCursor string, err error) {
	return nil, "", nil
}
