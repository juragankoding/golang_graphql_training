package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type categoriesRepository struct {
	Conn *sql.DB
}

func NewGenerateCategoriesRepository(Conn *sql.DB) domain.CategoriesRepository {
	return &categoriesRepository{
		Conn: Conn,
	}
}

func (r *categoriesRepository) Get(id int) (*domain.Categories, string, error) {
	return nil, "", nil
}

func (r *categoriesRepository) Fetch() (res []*domain.Categories, nextCursor string, err error) {
	return nil, "", nil
}

func (r *categoriesRepository) Insert(categories domain.Categories) (int64, string, error) {
	return -1, "", nil
}
func (r *categoriesRepository) Update(categories domain.Categories) (string, error) {
	return "", nil
}
func (r *categoriesRepository) Delete(id int) (string, error) {
	return "", nil
}
