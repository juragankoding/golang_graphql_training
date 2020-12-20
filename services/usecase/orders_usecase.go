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
	return o.ordersRepo.All()
}

func (o *ordersUseCase) Single(id int) (*domain.Orders, error) {
	return o.ordersRepo.Single(id)
}

func (o *ordersUseCase) Insert(orders domain.Orders) (int64, error) {
	return o.ordersRepo.Insert(orders)
}

func (o *ordersUseCase) Update(orders domain.Orders) (int64, error) {
	return o.ordersRepo.Update(orders)
}

func (o *ordersUseCase) Delete(id int) (int64, error) {
	return o.ordersRepo.Delete(id)
}
