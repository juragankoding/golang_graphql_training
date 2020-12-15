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

func (a *categoriesUseCase) Get(id int) (*domain.Categories, string, error) {
	return nil, "", nil
}

func (a *categoriesUseCase) Fetch() ([]*domain.Categories, string, error) {
	return nil, "", nil
}

func (a *categoriesUseCase) Insert(categories domain.Categories) (int64, string, error) {

	if categories.Name == "" {
		return -1, "name cannot null", errors.New("name cannot be null")
	}

	return -1, "", nil
}
func (a *categoriesUseCase) Update(categories domain.Categories) (string, error) {
	if err := categories.Validate(); err != nil {

	}

	return "", nil
}
func (a *categoriesUseCase) Delete(id int) (string, error) {
	return "", nil
}
