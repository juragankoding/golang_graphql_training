package usecase

import "github.com/juragankoding/golang_graphql_training/domain"

type stockUseCase struct {
	StockRepository domain.StocksRepository
}

func NewGenerateStockUseCase(stockRepository domain.StocksRepository) domain.StocksUseCase {
	return &stockUseCase{
		StockRepository: stockRepository,
	}
}

func (s *stockUseCase) All() ([]*domain.Stocks, error) {
	return nil, nil
}

func (s *stockUseCase) Single(id int) (*domain.Stocks, error) {
	return nil, nil
}

func (s *stockUseCase) Insert(stock domain.Stocks) (int64, error) {
	return -1, nil
}

func (s *stockUseCase) Update(stock domain.Stocks) (int64, error) {
	return -1, nil
}

func (s *stockUseCase) Delete(id int) (int64, error) {
	return -1, nil
}
