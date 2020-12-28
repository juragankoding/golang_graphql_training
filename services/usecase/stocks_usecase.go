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

func (s *stockUseCase) Fetch() ([]*domain.Stocks, error) {
	return s.StockRepository.Fetch()
}

func (s *stockUseCase) Get(stockID int, productID int) (*domain.Stocks, error) {
	return s.StockRepository.Get(stockID, productID)
}

func (s *stockUseCase) Insert(stock domain.Stocks) (int64, error) {
	return s.StockRepository.Insert(stock)
}

func (s *stockUseCase) Update(stock domain.Stocks) (int64, error) {
	return s.StockRepository.Update(stock)
}

func (s *stockUseCase) Delete(storeID int, productID int) (int64, error) {
	return s.StockRepository.Delete(storeID, productID)
}
