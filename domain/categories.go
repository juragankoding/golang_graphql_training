package domain

import "context"

type Categories struct {
	ID   int    `json:"category_id"`
	Name string `json:"category_name"`
}

type CategoriesUseCase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Categories, string, error)
}

type CategoriesRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Categories, nextCursor string, err error)
}
