package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var keyRedisOrderItem = "KeyRedisOrderItem"

type orderItemUseCase struct {
	orderItemRepository domain.OrderItemRepository
	contextUtil         context.Context
	redisUtil           *redis.Client
}

func NewGenerateOderItemUseCase(orderItemRepository domain.OrderItemRepository,
	contextUtil context.Context,
	redisUtil *redis.Client,
) domain.OrderItemUseCase {
	return &orderItemUseCase{
		orderItemRepository: orderItemRepository,
		contextUtil:         contextUtil,
		redisUtil:           redisUtil,
	}
}

func (o *orderItemUseCase) Fetch() ([]*domain.OrderItem, error) {
	var domainOrderItems []*domain.OrderItem

	var key = fmt.Sprint(keyRedisOrderItem, "_fetch")

	existsStatus, err := o.redisUtil.Exists(o.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsStatus <= 0 {
		domainOrderItems, err = o.orderItemRepository.Fetch()

		if err != nil {
			return nil, err
		}

		if jsonResult, err := json.Marshal(domainOrderItems); err != nil {
			return nil, err
		} else {
			if _, err := o.redisUtil.Set(o.contextUtil, key, jsonResult, cache.DurationFetch).Result(); err != nil {
				return nil, err
			}
		}
	} else {
		jsonFromRedis, err := o.redisUtil.Get(o.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), domainOrderItems); err != nil {
			return nil, err
		}
	}

	return domainOrderItems, nil
}

func (o *orderItemUseCase) Get(id int) (*domain.OrderItem, error) {

	var domainOrderItem *domain.OrderItem

	var key = fmt.Sprint(keyRedisOrderItem, "_get_", id)

	existsStatus, err := o.redisUtil.Exists(o.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsStatus <= 0 {
		domainOrderItem, err = o.orderItemRepository.Get(id)

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainOrderItem); err != nil {
			return nil, err
		} else {
			if _, err := o.redisUtil.Set(o.contextUtil, key, jsonDecode, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}
	} else {
		jsonFromRedis, err := o.redisUtil.Get(o.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainOrderItem); err != nil {
			return nil, err
		}
	}

	return domainOrderItem, nil
}

func (o *orderItemUseCase) Insert(orderItem domain.OrderItem) (int64, error) {
	return o.orderItemRepository.Insert(orderItem)
}

func (o *orderItemUseCase) Update(orderItem domain.OrderItem) (int64, error) {
	return o.orderItemRepository.Update(orderItem)
}

func (o *orderItemUseCase) Delete(id int) (int64, error) {
	return o.orderItemRepository.Delete(id)
}
