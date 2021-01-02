package usecase

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
	"github.com/juragankoding/golang_graphql_training/domain"
	"github.com/mattn/go/src/context"
)

var KeyRedisDatabaseCategories = "categories"

type categoriesUseCase struct {
	redisUtil      *redis.Client
	categoriesRepo domain.CategoriesRepository
	Context        context.Context
}

func NewGenerateCategoriesUserCase(a domain.CategoriesRepository,
	context context.Context,
	redisUtil *redis.Client) domain.CategoriesUseCase {
	return &categoriesUseCase{
		redisUtil:      redisUtil,
		categoriesRepo: a,
		Context:        context,
	}
}

func (a *categoriesUseCase) Get(id int) (*domain.Categories, error) {
	var domainCategory *domain.Categories

	key := fmt.Sprintf(KeyRedisDatabaseCategories, "_detail_", id)

	val, err := a.redisUtil.Exists(a.Context, key).Result()

	if err != nil {
		fmt.Printf("error redis: %s", err.Error())
	}

	if val <= 0 {
		domainCategory, err = a.categoriesRepo.Get(id)

		if err != nil {
			return nil, err
		}

		if jsonEncode, err := json.Marshal(domainCategory); err != nil {
			return nil, err
		} else {
			if _, err := a.redisUtil.Set(a.Context, key, jsonEncode, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}
	} else {
		result, err := a.redisUtil.Get(a.Context, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(result), &domainCategory); err != nil {
			return nil, err
		}
	}

	return domainCategory, nil
}

func (a *categoriesUseCase) Fetch() ([]*domain.Categories, error) {
	var domainCategories []*domain.Categories

	key := fmt.Sprint(KeyRedisDatabaseCategories, "_fetch")

	val, err := a.redisUtil.Exists(a.Context, key).Result()

	if err != nil {
		fmt.Printf("error redis : %s", err.Error())
	}

	if val <= 0 {
		domainCategories, err = a.categoriesRepo.Fetch()

		if err != nil {
			return nil, err
		}

		if e, err := json.Marshal(domainCategories); err != nil {
			return nil, err
		} else {
			_, err := a.redisUtil.Set(a.Context, key, e, cache.DurationFetch).Result()

			if err != nil {
				return nil, err
			}
		}
	} else {
		result, err := a.redisUtil.Get(a.Context, key).Result()

		if err != nil {
			return nil, err
		}

		erd := json.Unmarshal([]byte(result), &domainCategories)

		if erd != nil {
			return nil, err
		}
	}

	return domainCategories, nil
}

func (a *categoriesUseCase) Insert(categories *domain.Categories) (int64, error) {
	if categories.Name == "" {
		return -1, errors.New("name cannot be null")
	}

	return a.categoriesRepo.Insert(categories)
}
func (a *categoriesUseCase) Update(categories *domain.Categories) (int64, error) {
	if categories.Name == "" {
		return -1, errors.New("name cannot be null")
	}

	return a.categoriesRepo.Update(categories)
}
func (a *categoriesUseCase) Delete(id int) (int64, error) {
	if id >= 0 {
		return -1, errors.New("id must up from 0")
	}

	return a.categoriesRepo.Delete(id)
}
