package usecase

import (
	"github.com/juragankoding/golang_graphql_training/domain"
)

type ordersUseCase struct {
	ordersRepo domain.OrdersRepository
}

func NewGenerateOrdersUseCase(orderRepo domain.OrdersRepository) domain.OrdersUseCase {
	return &ordersUseCase{
		ordersRepo: orderRepo,
	}
}

func (o *ordersUseCase) All() ([]*domain.Orders, error) {
	return nil, nil
}

func (o *ordersUseCase) Single(id int) (*domain.Orders, error) {
	return nil, nil
}

func (o *ordersUseCase) Insert(orders domain.Orders) (int64, error) {
	return -1, nil
}

func (o *ordersUseCase) Update(orders domain.Orders) (int64, error) {
	return -1, nil
}

func (o *ordersUseCase) Delete(id int) (int64, error) {
	return -1, nil
}
