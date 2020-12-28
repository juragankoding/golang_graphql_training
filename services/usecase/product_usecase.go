package usecase

import (
	"github.com/juragankoding/golang_graphql_training/domain"
)

type productUsecase struct {
	productRepository domain.ProductsRepository
}

func NewGenerateProductUseCase(productRepository domain.ProductsRepository) domain.ProductsUseCase {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (a *productUsecase) Get(id int) (*domain.Products, error) {
	return a.productRepository.Get(id)
}

func (a *productUsecase) Fetch() ([]*domain.Products, error) {
	return a.productRepository.Fetch()
}

func (a *productUsecase) Insert(products domain.Products) (int64, error) {
	return a.productRepository.Insert(products)
}

func (a *productUsecase) Update(products domain.Products) (int64, error) {
	return a.productRepository.Update(products)
}

func (a *productUsecase) Delete(id int) (int64, error) {
	return a.productRepository.Delete(id)
}
