package domain

type Stocks struct {
	StoreID   int `json:"store_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type StocksRepository interface {
	All() ([]*Stocks, error)
	Single(id int) (*Stocks, error)
	Insert(stoks Stocks) (int64, error)
	Update(stocks Stocks) (int64, error)
	Delete(id int) (int64, error)
}

type StocksUseCase interface {
	All() ([]*Stocks, error)
	Single(id int) (*Stocks, error)
	Insert(stoks Stocks) (int64, error)
	Update(stocks Stocks) (int64, error)
	Delete(id int) (int64, error)
}
