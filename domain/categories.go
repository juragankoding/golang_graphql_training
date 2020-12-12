package domain

type Categories struct {
	ID   int    `json:"category_id"`
	Name string `json:"category_name"`
}

type CategoriesUseCase interface {
	Fetch() ([]Categories, string, error)
	Insert(Categories) (int64, string, error)
}

type CategoriesRepository interface {
	Fetch() (res []Categories, nextCursor string, err error)
	Insert(Categories) (int64, string, error)
}
