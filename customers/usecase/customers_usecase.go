package usecase

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type customersUseCase struct {
}

func NewCustomerRespository(Conn *sql.DB) domain.CustomersUseCase {
	return &customersUseCase{}
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
