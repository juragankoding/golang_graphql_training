package repository

import (
	"database/sql"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type stocksRepository struct {
	Conn *sql.DB
}

func NewGenerateStocksRepository(conn *sql.DB) domain.StocksRepository {
	return &stocksRepository{
		Conn: conn,
	}
}

func (s *stocksRepository) Fetch() ([]*domain.Stocks, error) {
	var listStock []*domain.Stocks

	query, err := s.Conn.Query("SELECT * FROM stocks")

	if err != nil {
		return listStock, err
	}

	for query.Next() {
		var stock domain.Stocks

		switch err := query.Scan(&stock.StoreID, &stock.ProductID, &stock.Quantity); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			listStock = append(listStock, &stock)
		}
	}

	return listStock, nil
}

func (s *stocksRepository) Get(storeID int, productID int) (*domain.Stocks, error) {

	fmt.Printf("SELECT * FROM stocks WHERE store_id=%d AND product_id=%d", storeID, productID)

	queryRow := s.Conn.QueryRow("SELECT * FROM stocks WHERE store_id=? AND product_id=?", storeID, productID)

	var stock domain.Stocks

	switch err := queryRow.Scan(&stock.StoreID, &stock.ProductID, &stock.Quantity); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return &stock, nil
	default:
		return nil, err
	}
}

func (s *stocksRepository) Insert(stocks domain.Stocks) (int64, error) {
	statement, err := s.Conn.Prepare("INSERT INTO stocks (store_id, product_id, quantity) VALUES (?,?,?)")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(stocks.StoreID, stocks.ProductID, stocks.Quantity); err != nil {
		return -1, err
	} else {
		return result.LastInsertId()
	}
}

func (s *stocksRepository) Update(stocks domain.Stocks) (int64, error) {
	statement, err := s.Conn.Prepare("UPDATE stocks SET quantity=? WHERE store_id = ? AND product_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(stocks.Quantity, stocks.StoreID, stocks.ProductID); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}

func (s *stocksRepository) Delete(storeID int, productID int) (int64, error) {
	statement, err := s.Conn.Prepare("DELETE FROM stocks WHERE store_id = ? AND product_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(storeID, productID); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}
