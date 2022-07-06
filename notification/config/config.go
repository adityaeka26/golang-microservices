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
			KafkaUrl:       os.Getenv("KAFKA_URL"),
			AppName:        os.Getenv("APP_NAME"),
			AppEnvironment: os.Getenv("APP_ENVIRONMENT"),
		},
	}
}

type Env struct {
	KafkaUrl       string
	AppName        string
	AppEnvironment string
}

func (config ConfigImpl) GetEnv() Env {
	return config.env
}
