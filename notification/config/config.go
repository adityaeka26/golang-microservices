package config

import "github.com/joho/godotenv"

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
		env: Env{},
	}
}

type Env struct{}

func (config ConfigImpl) GetEnv() Env {
	return config.env
}
