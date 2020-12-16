package usecase

import (
	"github.com/juragankoding/golang_graphql_training/domain"
)

type customersUseCase struct {
	customers domain.CustomersRepository
}

func NewGenerateCustomerRespository(customers domain.CustomersRepository) domain.CustomersUseCase {
	return &customersUseCase{
		customers: customers,
	}
}

func (a *customersUseCase) All() ([]*domain.Customers, error) {
	return nil, nil
}

func (a *customersUseCase) Get(id int) (*domain.Customers, error) {
	return nil, nil
}

func (a *customersUseCase) Insert(customer domain.Customers) (int64, error) {
	return -1, nil
}

func (a *customersUseCase) Update(customers domain.Customers) (int64, error) {
	return -1, nil
}

func (a *customersUseCase) Delete(id int) (int64, error) {
	return -1, nil
}
