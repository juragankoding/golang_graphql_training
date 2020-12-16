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

func (r *categoriesRepository) Get(id int) (res *domain.Categories, err error) {
	return nil, nil
}

func (r *categoriesRepository) Fetch() (res []*domain.Categories, err error) {
	return nil, nil
}

func (r *categoriesRepository) Insert(categories domain.Categories) (int64, error) {
	return -1, nil
}
func (r *categoriesRepository) Update(categories domain.Categories) (int64, error) {
	return -1, nil
}
func (r *categoriesRepository) Delete(id int) (int64, error) {
	return -1, nil
}
