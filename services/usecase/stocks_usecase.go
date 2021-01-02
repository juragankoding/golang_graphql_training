package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var keyRedisStocks = "keyRedisStocks"

type stockUseCase struct {
	StockRepository domain.StocksRepository
	contextUtil     context.Context
	redisUtil       *redis.Client
}

func NewGenerateStockUseCase(stockRepository domain.StocksRepository,
	contextUtil context.Context,
	redisUtil *redis.Client) domain.StocksUseCase {
	return &stockUseCase{
		StockRepository: stockRepository,
		contextUtil:     contextUtil,
		redisUtil:       redisUtil,
	}
}

func (s *stockUseCase) Fetch() ([]*domain.Stocks, error) {
	var domainStocks []*domain.Stocks

	var key = fmt.Sprint(keyRedisStocks, "_fetch_")

	existsResult, err := s.redisUtil.Exists(s.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsResult <= 0 {
		domainStocks, err = s.StockRepository.Fetch()

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainStocks); err != nil {
			return nil, err
		} else {
			if _, err := s.redisUtil.Set(s.contextUtil, key, jsonDecode, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}

	} else {
		jsonFromRedis, err := s.redisUtil.Get(s.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainStocks); err != nil {
			return nil, err
		}

	}

	return domainStocks, nil
}

func (s *stockUseCase) Get(stockID int, productID int) (*domain.Stocks, error) {
	var domainStocks *domain.Stocks

	var key = fmt.Sprint(keyRedisStocks, "_get_", stockID, "_", productID)

	existsResult, err := s.redisUtil.Exists(s.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsResult <= 0 {
		domainStocks, err = s.StockRepository.Get(stockID, productID)

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainStocks); err != nil {
			return nil, err
		} else {
			if _, err := s.redisUtil.Set(s.contextUtil, key, jsonDecode, cache.DurationDetail).Result(); err != nil {
				return nil, err
			}
		}

	} else {
		jsonFromRedis, err := s.redisUtil.Get(s.contextUtil, key).Result()

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainStocks); err != nil {
			return nil, err
		}

	}

	return domainStocks, nil
}

func (s *stockUseCase) Insert(stock domain.Stocks) (int64, error) {
	return s.StockRepository.Insert(stock)
}

func (s *stockUseCase) Update(stock domain.Stocks) (int64, error) {
	return s.StockRepository.Update(stock)
}

func (s *stockUseCase) Delete(storeID int, productID int) (int64, error) {
	return s.StockRepository.Delete(storeID, productID)
}
