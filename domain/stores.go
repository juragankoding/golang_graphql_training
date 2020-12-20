package domain

type Stores struct {
	StoreID   int    `json:"store_id"`
	StoreName string `json:"store_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
}

type StoresRepository interface {
	Single() (Stores, error)
	All() ([]Stores, error)
	Insert() (int64, error)
	Update() (int64, error)
	Delete() (int64, error)
}

type StoresUseCase interface {
	Single() (Stores, error)
	All() ([]Stores, error)
	Insert() (int64, error)
	Update() (int64, error)
	Delete() (int64, error)
}
