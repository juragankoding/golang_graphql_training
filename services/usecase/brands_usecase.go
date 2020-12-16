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
