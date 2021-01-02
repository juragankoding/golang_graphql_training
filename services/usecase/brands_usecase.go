package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var KeyRedisDatabaseBrands = "brands"

type brandsUseCase struct {
	brandsRepository domain.BrandsRepository
	Context          context.Context
	RedisUtil        *redis.Client
}

func NewGenerateBrandsUseCase(brandsRepository domain.BrandsRepository,
	contextUtil context.Context,
	redisUtil *redis.Client) domain.BrandsUseCase {

	return &brandsUseCase{
		brandsRepository: brandsRepository,
		Context:          contextUtil,
		RedisUtil:        redisUtil,
	}
}

func (b *brandsUseCase) Single(id int) (*domain.Brands, error) {
	var domainBrands *domain.Brands

	var key = fmt.Sprint(KeyRedisDatabaseBrands, "_detail_", id)

	result, err := b.RedisUtil.Exists(b.Context, key).Result()

	if err != nil {
		return nil, err
	}

	if result <= 0 {
		domainBrands, err = b.brandsRepository.Single(id)

		if err != nil {
			return nil, err
		}

		if jsonResult, err := json.Marshal(domainBrands); err != nil {
			return nil, err
		} else {
			if _, err := b.RedisUtil.Set(b.Context, key, jsonResult, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}

		}
	} else {
		jsonResult, err := b.RedisUtil.Get(b.Context, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonResult), &domainBrands); err != nil {
			return nil, err
		}
	}

	return domainBrands, nil
}

func (b *brandsUseCase) All() ([]*domain.Brands, error) {
	var domainFetch []*domain.Brands

	key := fmt.Sprintf(KeyRedisDatabaseBrands, "_fetch")

	result, err := b.RedisUtil.Exists(b.Context, key).Result()

	if err != nil {
		return nil, err
	}

	if result <= 0 {
		domainFetch, err = b.brandsRepository.All()

		if err != nil {
			return nil, err
		}

		if jsonResult, err := json.Marshal(domainFetch); err != nil {
			return nil, err
		} else {
			if _, err := b.RedisUtil.Set(b.Context, key, jsonResult, cache.DurationFetch).Result(); err != nil {
				return nil, err
			}
		}
	} else {
		jsonResult, err := b.RedisUtil.Get(b.Context, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonResult), &domainFetch); err != nil {
			return nil, err
		}
	}

	return domainFetch, nil
}

func (b *brandsUseCase) Insert(brands domain.Brands) (int64, error) {
	return b.brandsRepository.Insert(brands)
}

func (b *brandsUseCase) Update(brands domain.Brands) (int64, error) {
	return b.brandsRepository.Update(brands)
}

func (b *brandsUseCase) Delete(id int) (int64, error) {
	return b.brandsRepository.Delete(id)
}
