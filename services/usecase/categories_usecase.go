package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/db"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var KeyRedisDatabase = "categories"
var ctxBackground = context.Background()

type categoriesUseCase struct {
	redisUtil      *redis.Client
	categoriesRepo domain.CategoriesRepository
}

func NewGenerateCategoriesUserCase(a domain.CategoriesRepository) domain.CategoriesUseCase {
	return &categoriesUseCase{
		redisUtil:      db.GetRedisClient(),
		categoriesRepo: a,
	}
}

func (a *categoriesUseCase) Get(id int) (*domain.Categories, error) {
	var domainCategory *domain.Categories

	key := fmt.Sprintf(KeyRedisDatabase, "_detail_", id)

	val, err := a.redisUtil.Exists(ctxBackground, key).Result()

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
			if _, err := a.redisUtil.Set(ctxBackground, key, jsonEncode, db.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}
	} else {
		result, err := a.redisUtil.Get(ctxBackground, key).Result()

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

	key := fmt.Sprint(KeyRedisDatabase, "_fetch")

	val, err := a.redisUtil.Exists(ctxBackground, key).Result()

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
			_, err := a.redisUtil.Set(ctxBackground, key, e, db.DurationFetch).Result()

			if err != nil {
				return nil, err
			}
		}
	} else {
		result, err := a.redisUtil.Get(ctxBackground, key).Result()

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
