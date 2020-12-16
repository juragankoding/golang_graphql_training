package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type orderItemRepository struct {
	Conn *sql.DB
}

func NewGenerateOrderItemRespository(Conn *sql.DB) domain.OrderItemRepository {
	return &orderItemRepository{
		Conn: Conn,
	}
}

func (o *orderItemRepository) All() ([]*domain.OrderItem, error) {
	return nil, nil
}

func (o *orderItemRepository) Single(id int) (*domain.OrderItem, error) {
	return nil, nil
}

func (o *orderItemRepository) Insert(orderItem domain.OrderItem) (int64, error) {
	return -1, nil
}

func (o *orderItemRepository) Update(orderItem domain.OrderItem) (int64, error) {
	return -1, nil
}

func (o *orderItemRepository) Delete(id int) (int64, error) {
	return -1, nil
}
