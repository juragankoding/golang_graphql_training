package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var keyRedisStore = "keyRedisStore"

type storeUseCase struct {
	StoreRepo   *domain.StoresRepository
	contextUtil context.Context
	redisUtil   *redis.Client
}

func NewGenereateStoreUseCase(storeRepo *domain.StoresRepository,
	contextUtil context.Context,
	redisUtil *redis.Client) domain.StoresUseCase {
	store := storeUseCase{
		StoreRepo:   storeRepo,
		contextUtil: contextUtil,
		redisUtil:   redisUtil,
	}

	return &store
}

func (s *storeUseCase) Get(id int) (*domain.Stores, error) {
	var domainStores *domain.Stores

	var key = fmt.Sprint(keyRedisProduct, "_get_", id)

	existsResult, err := s.redisUtil.Exists(a.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsResult <= 0 {
		domainStores, err = s.StoreRepo.

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainStores); err != nil {
			return nil, err
		} else {
			if _, err := s.redisUtil.Set(a.contextUtil, key, jsonDecode, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}

	} else {
		jsonFromRedis, err := s.redisUtil.Get(a.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainStores); err != nil {
			return nil, err
		}

	}

	return domainStores, nil
}
func (s *storeUseCase) Fetch() ([]*domain.Stores, error) {
	var domainStores []*domain.Stores

	var key = fmt.Sprint(keyRedisProduct, "_get_", id)

	existsResult, err := s.redisUtil.Exists(a.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsResult <= 0 {
		domainStores, err = s.StoreRepo.

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainStores); err != nil {
			return nil, err
		} else {
			if _, err := s.redisUtil.Set(a.contextUtil, key, jsonDecode, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}

	} else {
		jsonFromRedis, err := s.redisUtil.Get(a.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainStores); err != nil {
			return nil, err
		}

	}

	return domainStores, nil
}
func (s *storeUseCase) Insert(stores domain.Stores) (int64, error) {
	return s.Insert(stores)
}
func (s *storeUseCase) Update(stores domain.Stores) (int64, error) {
	return s.Update(stores)
}
func (s *storeUseCase) Delete(id int) (int64, error) {
	return s.Delete(id)
}
