package domain

type Brands struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type BrandsUseCase interface {
	Single(id int) (*Brands, error)
	All() ([]*Brands, error)
	Insert(brands Brands) (int64, error)
	Update(brands Brands) (int64, error)
	Delete(id int) (int64, error)
}

type BrandsRepository interface {
	Single(id int) (*Brands, error)
	All() ([]*Brands, error)
	Insert(brands Brands) (int64, error)
	Update(brands Brands) (int64, error)
	Delete(id int) (int64, error)
}

func (brands *Brands) Compare(newBrands Brands) bool {
	return newBrands.ID == brands.ID &&
		brands.Name == newBrands.Name
}
