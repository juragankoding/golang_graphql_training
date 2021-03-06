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
	Fetch() ([]*OrderItem, error)
	Get(id int) (*OrderItem, error)
	Insert(orderItem OrderItem) (int64, error)
	Update(orderItem OrderItem) (int64, error)
	Delete(id int) (int64, error)
}

type OrderItemUseCase interface {
	Fetch() ([]*OrderItem, error)
	Get(id int) (*OrderItem, error)
	Insert(orderItem OrderItem) (int64, error)
	Update(orderItem OrderItem) (int64, error)
	Delete(id int) (int64, error)
}

func (o *OrderItem) Compare(orderItem OrderItem) bool {
	return o.ItemID == orderItem.ItemID &&
		o.Discount == orderItem.Discount &&
		o.ListPrice == orderItem.ListPrice &&
		o.Quantity == orderItem.Quantity &&
		o.OrderID == orderItem.OrderID &&
		o.ProductID == orderItem.ProductID
}
