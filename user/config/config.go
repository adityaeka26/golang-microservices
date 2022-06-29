package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	GetEnv() Env
}

type ConfigImpl struct {
	env Env
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return &ConfigImpl{
		env: Env{
			MongodbUrl:          os.Getenv("MONGODB_URL"),
			MongodbDatabaseName: os.Getenv("MONGODB_DATABASENAME"),
			JwtKey:              os.Getenv("JWT_KEY"),
			RedisUrl:            os.Getenv("REDIS_URL"),
			RedisPassword:       os.Getenv("REDIS_PASSWORD"),
		},
	}
}

type Env struct {
	MongodbUrl          string
	MongodbDatabaseName string
	JwtKey              string
	RedisUrl            string
	RedisPassword       string
}

func (config ConfigImpl) GetEnv() Env {
	return config.env
}
