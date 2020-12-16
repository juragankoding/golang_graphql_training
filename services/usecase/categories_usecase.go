package usecase

import (
	"errors"

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

func (a *categoriesUseCase) Get(id int) (*domain.Categories, error) {
	return nil, nil
}

func (a *categoriesUseCase) Fetch() ([]*domain.Categories, error) {
	return nil, nil
}

func (a *categoriesUseCase) Insert(categories domain.Categories) (int64, error) {

	if categories.Name == "" {
		return -1, errors.New("name cannot be null")
	}

	return -1, nil
}
func (a *categoriesUseCase) Update(categories domain.Categories) (int64, error) {
	if err := categories.Validate(); err != nil {

	}

	return -1, nil
}
func (a *categoriesUseCase) Delete(id int) (int64, error) {
	return -1, nil
}
