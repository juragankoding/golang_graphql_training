package usecase

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type productUsecase struct {
}

func NewProductUseCase(Conn sql.DB) domain.ProductsUseCase {
	return &productUsecase{}
}

func (a *productUsecase) Single() (domain.Products, error) {
	return domain.Products{}, nil
}

func (a *productUsecase) All() ([]domain.Products, error) {
	return nil, nil
}

func (a *productUsecase) Insert() (int64, error) {
	return -1, nil
}

func (a *productUsecase) Update() (int64, error) {
	return -1, nil
}

func (a *productUsecase) Delete() (int64, error) {
	return -1, nil
}
