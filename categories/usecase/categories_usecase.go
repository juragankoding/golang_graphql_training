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
<<<<<<< HEAD
	if err := categories.Validate(); err != nil {
		if lastID, _, err := a.categoriesRepo.Insert(categories); err != nil {
			return -1, err.Error(), err
		} else {
			return lastID, "Success insert data catgories", nil
		}

=======

	if categories.Name == "" {
		return -1, "name cannot null", errors.New("name cannot be null")
>>>>>>> 0f02cbf0ee4ebebd30fa40009ae30a26ad1bf6bb
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
