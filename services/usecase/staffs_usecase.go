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
	return s.StaffsRepo.All()
}

func (s *staffsUseCase) Single(id int) (*domain.Staffs, error) {
	return s.StaffsRepo.Single(id)
}

func (s *staffsUseCase) Insert(staffs domain.Staffs) (int64, error) {
	return s.StaffsRepo.Insert(staffs)
}

func (s *staffsUseCase) Update(staffs domain.Staffs) (int64, error) {
	return s.StaffsRepo.Update(staffs)
}

func (s *staffsUseCase) Delete(id int) (int64, error) {
	return s.StaffsRepo.Delete(id)
}
