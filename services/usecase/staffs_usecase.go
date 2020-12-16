package usecase

import (
	"github.com/juragankoding/golang_graphql_training/domain"
)

type staffsUseCase struct {
	StaffsRepo domain.StaffsRepository
}

func NewGenerateStaffsUseCase(staffRepo domain.StaffsRepository) domain.StaffsUseCase {
	return &staffsUseCase{
		StaffsRepo: staffRepo,
	}
}

func (s *staffsUseCase) All() ([]*domain.Staffs, error) {
	return nil, nil
}

func (s *staffsUseCase) Single(id int) (*domain.Staffs, error) {
	return nil, nil
}

func (s *staffsUseCase) Insert(staffs domain.Staffs) (int64, error) {
	return -1, nil
}

func (s *staffsUseCase) Update(staffs domain.Staffs) (int64, error) {
	return -1, nil
}

func (s *staffsUseCase) Delete(id int) (int64, error) {
	return -1, nil
}
