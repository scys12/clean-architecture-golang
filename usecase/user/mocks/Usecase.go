// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/scys12/clean-architecture-golang/model"
	mock "github.com/stretchr/testify/mock"

	request "github.com/scys12/clean-architecture-golang/pkg/payload/request"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// AuthenticateUser provides a mock function with given fields: _a0, _a1
func (_m *Usecase) AuthenticateUser(_a0 context.Context, _a1 *request.LoginRequest) (*model.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, *request.LoginRequest) *model.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *request.LoginRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditUserProfile provides a mock function with given fields: _a0
func (_m *Usecase) EditUserProfile(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterUser provides a mock function with given fields: _a0, _a1
func (_m *Usecase) RegisterUser(_a0 context.Context, _a1 *request.RegisterRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.RegisterRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
