package domain

type Categories struct {
	ID   int    `json:"category_id"`
	Name string `json:"category_name"`
}

type CategoriesUseCase interface {
	Fetch() ([]*Categories, string, error)
	Get(id int) (*Categories, string, error)
	Insert(categories Categories) (int64, string, error)
	Update(categories Categories) (string, error)
	Delete(id int) (string, error)
}

type CategoriesRepository interface {
	Fetch() (res []*Categories, nextCursor string, err error)
	Get(id int) (*Categories, string, error)
	Insert(categories Categories) (int64, string, error)
	Update(categories Categories) (string, error)
	Delete(id int) (string, error)
}

func (c *Categories) Validate() error {
	return nil
}
