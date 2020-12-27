package domain

import "database/sql"

type Customers struct {
	CustomerID int            `json:"customer_id"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	Phone      string         `json:"phone"`
	Email      string         `json:"email"`
	Street     string         `json:"street"`
	City       string         `json:"city"`
	State      sql.NullString `json:"state"`
	ZipCode    int            `json:"zip_code"`
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

func (c *Customers) Compare(customers Customers) bool {
	return c.CustomerID == customers.CustomerID &&
		c.FirstName == customers.FirstName &&
		c.LastName == customers.LastName &&
		c.Phone == customers.Phone &&
		c.Email == customers.Email &&
		c.Street == customers.Street &&
		c.City == customers.City &&
		// c.State == customers.State &&
		c.ZipCode == customers.ZipCode
}
