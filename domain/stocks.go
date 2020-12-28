package domain

type Stocks struct {
	StoreID   int `json:"store_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type StocksRepository interface {
	Fetch() ([]*Stocks, error)
	Get(storeID int, productID int) (*Stocks, error)
	Insert(stoks Stocks) (int64, error)
	Update(stocks Stocks) (int64, error)
	Delete(storeID int, productID int) (int64, error)
}

type StocksUseCase interface {
	Fetch() ([]*Stocks, error)
	Get(storeID int, productID int) (*Stocks, error)
	Insert(stoks Stocks) (int64, error)
	Update(stocks Stocks) (int64, error)
	Delete(storeID int, productID int) (int64, error)
}

func (s *Stocks) Compare(stock Stocks) bool {
	return s.StoreID == stock.StoreID &&
		s.ProductID == stock.ProductID &&
		s.Quantity == stock.Quantity
}
