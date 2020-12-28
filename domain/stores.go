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
	Get(id int) (*Stores, error)
	Fetch() ([]*Stores, error)
	Insert(stores Stores) (int64, error)
	Update(stores Stores) (int64, error)
	Delete(id int) (int64, error)
}

type StoresUseCase interface {
	Get(id int) (*Stores, error)
	Fetch() ([]*Stores, error)
	Insert(stores Stores) (int64, error)
	Update(stores Stores) (int64, error)
	Delete(id int) (int64, error)
}

func (s *Stores) Compare(stores Stores) bool {
	return s.StoreID == stores.StoreID &&
		s.StoreName == stores.StoreName &&
		s.Email == stores.Email &&
		s.City == stores.City &&
		s.Phone == stores.Phone &&
		s.State == stores.State &&
		s.Street == stores.Street &&
		s.ZipCode == stores.ZipCode
}
