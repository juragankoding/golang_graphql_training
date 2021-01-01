package db

import "github.com/go-redis/redis/v8"

func GetRedisClient() *redis.Client {
	edb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return edb
}
