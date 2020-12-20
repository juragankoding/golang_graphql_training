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

func (a *productUsecase) Single(id int) (*domain.Products, error) {
	return a.productRepository.Single(id)
}

func (a *productUsecase) All() ([]*domain.Products, error) {
	return a.productRepository.All()
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
