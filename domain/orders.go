package domain

type Orders struct {
	OrderID      int    `json:"order_id"`
	CustomerID   int    `json:"customer_id"`
	OrderStatus  int    `json:"order_status"`
	OrderDate    string `json:"order_date"`
	RequiredDate string `json:"required_date"`
	ShippedDate  string `json:"shipped_date"`
	StoreID      int    `json:"store_id"`
	StaffID      int    `json:"staff_id"`
}

type OrdersRepository interface {
	Fetch() ([]*Orders, error)
	Get(id int) (*Orders, error)
	Insert(order Orders) (int64, error)
	Update(order Orders) (int64, error)
	Delete(id int) (int64, error)
}

type OrdersUseCase interface {
	Fetch() ([]*Orders, error)
	Get(id int) (*Orders, error)
	Insert(orders Orders) (int64, error)
	Update(orders Orders) (int64, error)
	Delete(id int) (int64, error)
}

func (o *Orders) Compare(orders Orders) bool {
	return o.OrderID == orders.OrderID &&
		// o.CustomerID == orders.CustomerID &&
		// o.OrderStatus == orders.OrderStatus &&
		// o.OrderDate == orders.OrderDate &&
		// o.RequiredDate == orders.RequiredDate &&
		// o.ShippedDate == orders.ShippedDate &&
		// o.StoreID == orders.StoreID &&
		o.StaffID == orders.StaffID
}
