package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var keyRedisOrders = "keyRedisOrders"

type ordersUseCase struct {
	ordersRepo  domain.OrdersRepository
	contextUtil context.Context
	redisUtil   *redis.Client
}

func NewGenerateOrdersUseCase(orderRepo domain.OrdersRepository,
	contextUtil context.Context,
	redisUtil *redis.Client) domain.OrdersUseCase {
	return &ordersUseCase{
		ordersRepo:  orderRepo,
		contextUtil: contextUtil,
		redisUtil:   redisUtil,
	}
}

func (o *ordersUseCase) Fetch() ([]*domain.Orders, error) {
	var domainOrder []*domain.Orders

	var key = fmt.Sprint(keyRedisOrders, "_fetch")

	result, err := o.redisUtil.Exists(o.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if result <= 0 {
		domainOrder, err = o.ordersRepo.Fetch()

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainOrder); err != nil {
			return nil, err
		} else {
			if _, err := o.redisUtil.Set(o.contextUtil, key, jsonDecode, cache.DurationFetch).Result(); err != nil {
				return nil, err
			}
		}
	} else {
		jsonFromRedis, err := o.redisUtil.Get(o.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainOrder); err != nil {
			return nil, err
		}
	}

	return domainOrder, err
}

func (o *ordersUseCase) Get(id int) (*domain.Orders, error) {
	var domainOrders *domain.Orders

	var key = fmt.Sprint(keyRedisOrders, "_get_", id)

	resultExists, err := o.redisUtil.Exists(o.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if resultExists <= 0 {
		domainOrders, err := o.ordersRepo.Get(id)

		if err != nil {
			return nil, err
		}

		if jsonMarshal, err := json.Marshal(domainOrders); err != nil {
			return nil, err
		} else {
			if _, err := o.redisUtil.Set(o.contextUtil, key, jsonMarshal, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}
	} else {
		jsonFromRedis, err := o.redisUtil.Get(o.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), domainOrders); err != nil {
			return nil, err
		}
	}

	return domainOrders, nil
}

func (o *ordersUseCase) Insert(orders domain.Orders) (int64, error) {
	return o.ordersRepo.Insert(orders)
}

func (o *ordersUseCase) Update(orders domain.Orders) (int64, error) {
	return o.ordersRepo.Update(orders)
}

func (o *ordersUseCase) Delete(id int) (int64, error) {
	return o.ordersRepo.Delete(id)
}
