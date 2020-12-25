package usecase

import (
	"errors"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type categoriesUseCase struct {
	categoriesRepo domain.CategoriesRepository
}

func NewGenerateCategoriesUserCase(a domain.CategoriesRepository) domain.CategoriesUseCase {
	return &categoriesUseCase{
		categoriesRepo: a,
	}
}

func (a *categoriesUseCase) Get(id int) (*domain.Categories, error) {
	return a.categoriesRepo.Get(id)
}

func (a *categoriesUseCase) Fetch() ([]*domain.Categories, error) {
	return a.categoriesRepo.Fetch()
}

func (a *categoriesUseCase) Insert(categories *domain.Categories) (int64, error) {
	if categories.Name == "" {
		return -1, errors.New("name cannot be null")
	}

	return a.categoriesRepo.Insert(categories)
}
func (a *categoriesUseCase) Update(categories *domain.Categories) (int64, error) {
	if categories.Name == "" {
		return -1, errors.New("name cannot be null")
	}

	return a.categoriesRepo.Update(categories)
}
func (a *categoriesUseCase) Delete(id int) (int64, error) {
	if id >= 0 {
		return -1, errors.New("id must up from 0")
	}

	return a.categoriesRepo.Delete(id)
}
