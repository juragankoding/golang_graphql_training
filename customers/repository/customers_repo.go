package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type customerRepository struct {
	Conn *sql.DB
}

func NewCustomerRespository(Conn *sql.DB) domain.CustomersRepository {
	return &customerRepository{
		Conn: Conn,
	}
}

func (a *customerRepository) All() ([]*domain.Customers, error) {
	return nil, nil
}

func (a *customerRepository) Get(id int) (*domain.Customers, error) {
	return nil, nil
}

func (a *customerRepository) Insert(customer domain.Customers) (int64, error) {
	return -1, nil
}

func (a *customerRepository) Update(customers domain.Customers) (int64, error) {
	return -1, nil
}

func (a *customerRepository) Delete(id int) (int64, error) {
	return -1, nil
}
