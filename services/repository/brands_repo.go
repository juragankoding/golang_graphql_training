package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type brandsRepository struct {
	Conn *sql.DB
}

func NewGenerateBrandsRepository(Conn *sql.DB) domain.BrandsRepository {
	return &brandsRepository{
		Conn: Conn,
	}
}

func (b *brandsRepository) Single() (domain.Brands, error) {
	return domain.Brands{
		ID:   0,
		Name: "",
	}, nil
}

func (b *brandsRepository) All() ([]domain.Brands, error) {
	return nil, nil
}

func (b *brandsRepository) Insert() (int64, error) {
	return -1, nil
}

func (b *brandsRepository) Update() (int64, error) {
	return -1, nil
}

func (b *brandsRepository) Delete() (int64, error) {
	return -1, nil
}
