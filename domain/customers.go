package domain

type Customers struct {
	CustomerID int    `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    int    `json:"zip_code"`
}

type CustomersUseCase interface {
	All() (res []*Customers, err error)
	Get(id int) (*Customers, error)
	Insert(customers Customers) (int64, error)
	Update(customers Customers) (int64, error)
	Delete(id int) (int64, error)
}

type CustomersRepository interface {
	All() (res []*Customers, err error)
	Get(id int) (*Customers, error)
	Insert(customers Customers) (int64, error)
	Update(customers Customers) (int64, error)
	Delete(id int) (int64, error)
}

func (c *Customers) Validate() error {
	return nil
}
