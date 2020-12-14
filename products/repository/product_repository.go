package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type productsRepository struct {
	Conn *sql.DB
}

func NewProductsRepostory(Conn *sql.DB) domain.ProductsRepository {
	return &productsRepository{
		Conn: Conn,
	}
}

func (a *productsRepository) Single() (domain.Products, error) {
	return domain.Products{}, nil
}

func (a *productsRepository) All() ([]domain.Products, error) {
	return nil, nil
}

func (a *productsRepository) Insert() (int64, error) {
	return -1, nil
}

func (a *productsRepository) Update() (int64, error) {
	return -1, nil
}

func (a *productsRepository) Delete() (int64, error) {
	return -1, nil
}
