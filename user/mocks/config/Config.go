// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	config "github.com/adityaeka26/golang-microservices/user/config"
	mock "github.com/stretchr/testify/mock"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// GetEnv provides a mock function with given fields:
func (_m *Config) GetEnv() config.Env {
	ret := _m.Called()

	var r0 config.Env
	if rf, ok := ret.Get(0).(func() config.Env); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.Env)
	}

	return r0
}

type mockConstructorTestingTNewConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfig creates a new instance of Config. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfig(t mockConstructorTestingTNewConfig) *Config {
	mock := &Config{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
