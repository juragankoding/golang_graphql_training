package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var keyRedisProduct = "keyRedisProduct"

type productUsecase struct {
	productRepository domain.ProductsRepository
	contextUtil       context.Context
	redisUtil         *redis.Client
}

func NewGenerateProductUseCase(productRepository domain.ProductsRepository,
	contextUtil context.Context,
	redisUtil *redis.Client) domain.ProductsUseCase {
	return &productUsecase{
		productRepository: productRepository,
		contextUtil:       contextUtil,
		redisUtil:         redisUtil,
	}
}

func (a *productUsecase) Get(id int) (*domain.Products, error) {
	var domainProduct *domain.Products

	var key = fmt.Sprint(keyRedisProduct, "_get_", id)

	existsResult, err := a.redisUtil.Exists(a.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsResult <= 0 {
		domainProduct, err = a.productRepository.Get(id)

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainProduct); err != nil {
			return nil, err
		} else {
			if _, err := a.redisUtil.Set(a.contextUtil, key, jsonDecode, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}

	} else {
		jsonFromRedis, err := a.redisUtil.Get(a.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainProduct); err != nil {
			return nil, err
		}

	}

	return domainProduct, nil
}

func (a *productUsecase) Fetch() ([]*domain.Products, error) {
	var domainProduct []*domain.Products

	var key = fmt.Sprint(keyRedisProduct, "_fetch")

	existsResult, err := a.redisUtil.Exists(a.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsResult <= 0 {
		domainProduct, err = a.productRepository.Fetch()

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainProduct); err != nil {
			return nil, err
		} else {
			if _, err := a.redisUtil.Set(a.contextUtil, key, jsonDecode, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}

	} else {
		jsonFromRedis, err := a.redisUtil.Get(a.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainProduct); err != nil {
			return nil, err
		}

	}

	return domainProduct, nil
}

func (a *productUsecase) Insert(products domain.Products) (int64, error) {
	return a.productRepository.Insert(products)
}

func (a *productUsecase) Update(products domain.Products) (int64, error) {
	return a.productRepository.Update(products)
}

func (a *productUsecase) Delete(id int) (int64, error) {
	return a.productRepository.Delete(id)
}
