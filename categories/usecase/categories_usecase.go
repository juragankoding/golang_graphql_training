package usecase

import (
	"context"

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

func (a *categoriesUseCase) Fetch(ctx context.Context, cursor string, num int64) ([]domain.Categories, string, error) {
	return nil, "", nil
}
