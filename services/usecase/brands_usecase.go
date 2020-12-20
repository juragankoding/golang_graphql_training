package usecase

import (
	"github.com/juragankoding/golang_graphql_training/domain"
)

type brandsUseCase struct {
	brandsRepository domain.BrandsRepository
}

func NewGenerateBrandsUseCase(brandsRepository domain.BrandsRepository) domain.BrandsUseCase {
	return &brandsUseCase{
		brandsRepository: brandsRepository,
	}
}

func (b *brandsUseCase) Single(id int) (*domain.Brands, error) {
	return b.brandsRepository.Single(id)
}

func (b *brandsUseCase) All() ([]*domain.Brands, error) {
	return b.brandsRepository.All()
}

func (b *brandsUseCase) Insert(brands domain.Brands) (int64, error) {
	return b.brandsRepository.Insert(brands)
}

func (b *brandsUseCase) Update(brands domain.Brands) (int64, error) {
	return b.brandsRepository.Update(brands)
}

func (b *brandsUseCase) Delete(id int) (int64, error) {
	return b.brandsRepository.Delete(id)
}
