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
	return s.StockRepository.All()
}

func (s *stockUseCase) Single(id int) (*domain.Stocks, error) {
	return s.StockRepository.Single(id)
}

func (s *stockUseCase) Insert(stock domain.Stocks) (int64, error) {
	return s.StockRepository.Insert(stock)
}

func (s *stockUseCase) Update(stock domain.Stocks) (int64, error) {
	return s.StockRepository.Update(stock)
}

func (s *stockUseCase) Delete(id int) (int64, error) {
	return s.StockRepository.Delete(id)
}
