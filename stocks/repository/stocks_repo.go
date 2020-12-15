package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type stocksRepository struct {
	Conn *sql.DB
}

func NewStocksRepository(Conn *sql.DB) domain.StocksRepository {
	return &stocksRepository{
		Conn: Conn,
	}
}

func (s *stocksRepository) All() ([]*domain.Stocks, error) {
	return nil, nil
}

func (s *stocksRepository) Single(id int) (*domain.Stocks, error) {
	return nil, nil
}

func (s *stocksRepository) Update(stocks domain.Stocks) (int64, error) {
	return -1, nil
}

func (s *stocksRepository) Insert(stocks domain.Stocks) (int64, error) {
	return -1, nil
}

func (s *stocksRepository) Delete(id int) (int64, error) {
	return -1, nil
}
