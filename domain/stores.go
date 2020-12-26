package domain

type Stores struct {
	StoreID   int    `json:"store_id"`
	StoreName string `json:"store_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
}

type StoresRepository interface {
	Single(id int) (*Stores, error)
	All() ([]*Stores, error)
	Insert(stores Stores) (int64, error)
	Update(stores Stores) (int64, error)
	Delete(id int) (int64, error)
}

type StoresUseCase interface {
	Single(id int) (*Stores, error)
	All() ([]*Stores, error)
	Insert(stores Stores) (int64, error)
	Update(stores Stores) (int64, error)
	Delete(id int) (int64, error)
}
