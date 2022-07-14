// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	zap "go.uber.org/zap"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// GetLogger provides a mock function with given fields:
func (_m *Logger) GetLogger() *zap.Logger {
	ret := _m.Called()

	var r0 *zap.Logger
	if rf, ok := ret.Get(0).(func() *zap.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zap.Logger)
		}
	}

	return r0
}

type mockConstructorTestingTNewLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogger(t mockConstructorTestingTNewLogger) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}