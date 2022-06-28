package database

import "github.com/go-redis/redis/v8"

type Redis interface {
	GetClient() *redis.Client
}
