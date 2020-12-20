package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type productsRepository struct {
	Conn *sql.DB
}

func NewGenerateProductsRepository(Conn *sql.DB) domain.ProductsRepository {
	return &productsRepository{
		Conn: Conn,
	}
}

func (a *productsRepository) Single(id int) (*domain.Products, error) {
	return nil, nil
}

func (a *productsRepository) All() ([]*domain.Products, error) {
	return nil, nil
}

func (a *productsRepository) Insert(products domain.Products) (int64, error) {
	return -1, nil
}

func (a *productsRepository) Update(products domain.Products) (int64, error) {
	return -1, nil
}

func (a *productsRepository) Delete(id int) (int64, error) {
	return -1, nil
}
