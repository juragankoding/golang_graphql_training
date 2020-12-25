package domain

type Categories struct {
	ID   int    `json:"category_id"`
	Name string `json:"category_name"`
}

type CategoriesUseCase interface {
	Fetch() ([]*Categories, error)
	Get(id int) (*Categories, error)
	Insert(categories *Categories) (int64, error)
	Update(categories *Categories) (int64, error)
	Delete(id int) (int64, error)
}

type CategoriesRepository interface {
	Fetch() (res []*Categories, err error)
	Get(id int) (*Categories, error)
	Insert(categories *Categories) (int64, error)
	Update(categories *Categories) (int64, error)
	Delete(id int) (int64, error)
}

func (c *Categories) Validate() error {
	return nil
}
