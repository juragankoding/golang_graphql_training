package domain

type Products struct {
	ProductID      int        `json:"product_id"`
	ProductName    string     `json:"product_name"`
	BrandID        int        `json:"brand_id"`
	CategoryID     int        `json:"category_id"`
	ModelYear      string     `json:"model_year"`
	ListPrice      int        `json:"list_price"`
	CategoriesData Categories `json:"categories_data"`
	BrandsData     Brands     `json:"brands_data"`
}

type ProductsRepository interface {
	Single(id int) (*Products, error)
	All() ([]*Products, error)
	Insert(products Products) (int64, error)
	Update(products Products) (int64, error)
	Delete(id int) (int64, error)
}

type ProductsUseCase interface {
	Single(id int) (*Products, error)
	All() ([]*Products, error)
	Insert(products Products) (int64, error)
	Update(products Products) (int64, error)
	Delete(id int) (int64, error)
}
