package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type ordersRepository struct {
	Conn *sql.DB
}

func NewOrdersRepository(Conn *sql.DB) domain.OrdersRepository {
	return &ordersRepository{
		Conn: Conn,
	}
}

func (o *ordersRepository) All() ([]*domain.Orders, error) {
	return nil, nil
}

func (o *ordersRepository) Single(id int) (*domain.Orders, error) {
	return nil, nil
}

func (o *ordersRepository) Insert(orders domain.Orders) (int64, error) {
	return -1, nil
}

func (o *ordersRepository) Update(orders domain.Orders) (int64, error) {
	return -1, nil
}
func (o *ordersRepository) Delete(id int) (int64, error) {
	return -1, nil
}
