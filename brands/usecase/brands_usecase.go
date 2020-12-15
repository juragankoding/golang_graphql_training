package usecase

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type brandsUseCase struct {
	Conn *sql.DB
}

func NewBrandsRepository(Conn *sql.DB) domain.BrandsUseCase {
	return &brandsUseCase{
		Conn: Conn,
	}
}

func (b *brandsUseCase) Single() (domain.Brands, error) {
	return domain.Brands{
		ID:   0,
		Name: "",
	}, nil
}

func (b *brandsUseCase) All() ([]domain.Brands, error) {
	return nil, nil
}

func (b *brandsUseCase) Insert() (int64, error) {
	return -1, nil
}

func (b *brandsUseCase) Update() (int64, error) {
	return -1, nil
}

func (b *brandsUseCase) Delete() (int64, error) {
	return -1, nil
}
