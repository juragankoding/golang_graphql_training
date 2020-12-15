package domain

type Brands struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type BrandsUseCase interface {
	Single() (Brands, error)
	All() ([]Brands, error)
	Insert() (int64, error)
	Update() (int64, error)
	Delete() (int64, error)
}

type BrandsRepository interface {
	Single() (Brands, error)
	All() ([]Brands, error)
	Insert() (int64, error)
	Update() (int64, error)
	Delete() (int64, error)
}
