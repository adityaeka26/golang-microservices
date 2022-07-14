// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/adityaeka26/golang-microservices/user/module/model/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindOneUser provides a mock function with given fields: ctx, filter
func (_m *Repository) FindOneUser(ctx context.Context, filter interface{}) (*domain.User, error) {
	ret := _m.Called(ctx, filter)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *domain.User); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOneUser provides a mock function with given fields: ctx, document
func (_m *Repository) InsertOneUser(ctx context.Context, document interface{}) (*string, error) {
	ret := _m.Called(ctx, document)

	var r0 *string
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *string); ok {
		r0 = rf(ctx, document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}