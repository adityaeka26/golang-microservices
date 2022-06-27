package config

type Config interface {
	GetEnv() Env
}
