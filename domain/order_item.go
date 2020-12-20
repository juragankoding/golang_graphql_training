package domain

type OrderItem struct {
	ItemID    int `json:"item_id"`
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	ListPrice int `json:"list_price"`
	Discount  int `json:"discount"`
}

type OrderItemRepository interface {
	All() ([]*OrderItem, error)
	Single(id int) (*OrderItem, error)
	Insert(orderItem OrderItem) (int64, error)
	Update(orderItem OrderItem) (int64, error)
	Delete(id int) (int64, error)
}

type OrderItemUseCase interface {
	All() ([]*OrderItem, error)
	Single(id int) (*OrderItem, error)
	Insert(orderItem OrderItem) (int64, error)
	Update(orderItem OrderItem) (int64, error)
	Delete(id int) (int64, error)
}
