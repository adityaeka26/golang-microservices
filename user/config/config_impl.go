package config

import (
	"os"

	"github.com/joho/godotenv"
)

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
		},
	}
}

type Env struct {
	MongodbUrl          string
	MongodbDatabaseName string
}

func (config ConfigImpl) GetEnv() Env {
	return config.env
}
