package usecase

import (
	"github.com/juragankoding/golang_graphql_training/domain"
)

type storeUseCase struct {
	StoreRepo *domain.StoresRepository
}

func NewGenereateStoreUseCase(storeRepo *domain.StoresRepository) domain.StoresUseCase {
	store := storeUseCase{
		StoreRepo: storeRepo,
	}

	return &store
}

func (s *storeUseCase) Single(id int) (*domain.Stores, error) {
	return s.Single(id)
}
func (s *storeUseCase) All() ([]*domain.Stores, error) {
	return s.All()
}
func (s *storeUseCase) Insert(stores domain.Stores) (int64, error) {
	return s.Insert(stores)
}
func (s *storeUseCase) Update(stores domain.Stores) (int64, error) {
	return s.Update(stores)
}
func (s *storeUseCase) Delete(id int) (int64, error) {
	return s.Delete(id)
}
