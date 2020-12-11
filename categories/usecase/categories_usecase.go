package usecase

import (
	"github.com/juragankoding/golang_graphql_training/domain"
)

type categoriesUseCase struct {
	categoriesRepo domain.CategoriesRepository
}

func NewCategoriesUserCase(a domain.CategoriesRepository) domain.CategoriesUseCase {
	return &categoriesUseCase{
		categoriesRepo: a,
	}
}

func (a *categoriesUseCase) Fetch() ([]domain.Categories, string, error) {
	return nil, "", nil
}

func (a *categoriesUseCase) Insert(categories domain.Categories) (int64, string, error) {
	return -1, "", nil
}
func (a *categoriesUseCase) Update(categories domain.Categories) (int64, string, error) {
	return -1, "", nil
}
func (a *categoriesUseCase) Delete(categories domain.Categories) (int64, string, error) {
	return -1, "", nil
}
