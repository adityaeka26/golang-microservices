package database

import "github.com/go-redis/redis/v8"

type RedisImpl struct {
	redisClient *redis.Client
}

func NewRedis() Redis {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "leomessi",
		DB:       0,
	})

	return &RedisImpl{
		redisClient: redisClient,
	}
}

func (redis RedisImpl) GetClient() *redis.Client {
	return redis.redisClient
}
