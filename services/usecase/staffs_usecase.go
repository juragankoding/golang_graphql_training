package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/juragankoding/golang_graphql_training/cache"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var keyRedisStaff = "keyRedisStaff"

type staffsUseCase struct {
	StaffsRepo  domain.StaffsRepository
	contextUtil context.Context
	redisUtil   *redis.Client
}

func NewGenerateStaffsUseCase(staffRepo domain.StaffsRepository,
	contextUtil context.Context,
	redisUtil *redis.Client) domain.StaffsUseCase {
	return &staffsUseCase{
		StaffsRepo:  staffRepo,
		contextUtil: contextUtil,
		redisUtil:   redisUtil,
	}
}

func (s *staffsUseCase) All() ([]*domain.Staffs, error) {
	var domainStaffs []*domain.Staffs

	var key = fmt.Sprint(keyRedisStaff, "_fetch_")

	existsResult, err := s.redisUtil.Exists(s.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsResult <= 0 {
		domainStaffs, err = s.StaffsRepo.All()

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainStaffs); err != nil {
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

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainStaffs); err != nil {
			return nil, err
		}

	}

	return domainStaffs, nil
}

func (s *staffsUseCase) Single(id int) (*domain.Staffs, error) {
	var domainStaffs *domain.Staffs

	var key = fmt.Sprint(keyRedisStaff, "_get_", id)

	existsResult, err := s.redisUtil.Exists(s.contextUtil, key).Result()

	if err != nil {
		return nil, err
	}

	if existsResult <= 0 {
		domainStaffs, err = s.StaffsRepo.Single(id)

		if err != nil {
			return nil, err
		}

		if jsonDecode, err := json.Marshal(domainStaffs); err != nil {
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

		if err := json.Unmarshal([]byte(jsonFromRedis), &domainStaffs); err != nil {
			return nil, err
		}

	}

	return domainStaffs, nil
}

func (s *staffsUseCase) Insert(staffs domain.Staffs) (int64, error) {
	if err := staffs.Validate(); err != nil {
		return -1, err
	}

	status, err := s.StaffsRepo.Insert(staffs)

	return status, err
}

func (s *staffsUseCase) Update(staffs domain.Staffs) (int64, error) {
	if err := staffs.Validate(); err != nil {
		return -1, err
	}

	return s.StaffsRepo.Update(staffs)
}

func (s *staffsUseCase) Delete(id int) (int64, error) {
	return s.StaffsRepo.Delete(id)
}
