package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var keyRedisCustomers = "KeyRedisCustomer"

type customersUseCase struct {
	customers   domain.CustomersRepository
	ContextUtil context.Context
	RedisUtil   *redis.Client
}

func NewGenerateCustomerUseCase(customers domain.CustomersRepository,
	contextUtil context.Context,
	redisUtil *redis.Client,
) domain.CustomersUseCase {
	return &customersUseCase{
		customers:   customers,
		ContextUtil: contextUtil,
		RedisUtil:   redisUtil,
	}
}

func (a *customersUseCase) All() ([]*domain.Customers, error) {
	var domainCustomers []*domain.Customers

	var key = fmt.Sprint(keyRedisCustomers, "_fetch")

	result, err := a.RedisUtil.Exists(a.ContextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if result <= 0 {
		domainCustomers, err = a.customers.All()

		if err != nil {
			return nil, err
		}

		if resultJson, err := json.Marshal(domainCustomers); err != nil {
			return nil, err
		} else {
			if err := a.RedisUtil.Set(a.ContextUtil, key, resultJson, cache.DurationFetch); err != nil {
				fmt.Printf("redis find error: %s", err.String())
			}
		}
	} else {
		resultJson, err := a.RedisUtil.Get(a.ContextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(resultJson), &resultJson); err != nil {
			return nil, err
		}
	}

	return domainCustomers, nil
}

func (a *customersUseCase) Get(id int) (*domain.Customers, error) {
	var domainCustomer *domain.Customers

	var key = fmt.Sprint(keyRedisCustomers, "_get_", id)

	result, err := a.RedisUtil.Exists(a.ContextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if result <= 0 {
		domainCustomer, err = a.customers.Get(id)

		if err != nil {
			return nil, err
		}

		if resultJson, err := json.Marshal(domainCustomer); err != nil {
			return nil, err
		} else {
			if _, err := a.RedisUtil.Set(a.ContextUtil, key, resultJson, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}
	} else {
		resultJsonFromRedis, err := a.RedisUtil.Get(a.ContextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(resultJsonFromRedis), &domainCustomer); err != nil {
			return nil, err
		}
	}

	return domainCustomer, nil
}

func (a *customersUseCase) Insert(customer domain.Customers) (int64, error) {
	return a.customers.Insert(customer)
}

func (a *customersUseCase) Update(customers domain.Customers) (int64, error) {
	return a.customers.Update(customers)
}

func (a *customersUseCase) Delete(id int) (int64, error) {
	return a.customers.Delete(id)
}
