package database

import "github.com/go-redis/redis/v8"

type Redis interface {
	GetClient() *redis.Client
}
type RedisImpl struct {
	redisClient *redis.Client
}

func NewRedis(addr string, password string) Redis {
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
