package usecase

import "github.com/juragankoding/golang_graphql_training/domain"

type orderItemUseCase struct {
	orderItemRepository domain.OrderItemRepository
}

func NewGenerateOderItemUseCase(orderItemRepository domain.OrderItemRepository) domain.OrderItemUseCase {
	return &orderItemUseCase{
		orderItemRepository: orderItemRepository,
	}
}

func (o *orderItemUseCase) All() ([]*domain.OrderItem, error) {
	return o.orderItemRepository.All()
}

func (o *orderItemUseCase) Single(id int) (*domain.OrderItem, error) {
	return o.orderItemRepository.Single(id)
}

func (o *orderItemUseCase) Insert(orderItem domain.OrderItem) (int64, error) {
	return o.orderItemRepository.Insert(orderItem)
}

func (o *orderItemUseCase) Update(orderItem domain.OrderItem) (int64, error) {
	return o.orderItemRepository.Update(orderItem)
}

func (o *orderItemUseCase) Delete(id int) (int64, error) {
	return o.orderItemRepository.Delete(id)
}
