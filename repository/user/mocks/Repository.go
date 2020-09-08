// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/scys12/clean-architecture-golang/model"
	mock "github.com/stretchr/testify/mock"

	request "github.com/scys12/clean-architecture-golang/payload/request"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetUserAuthenticateData provides a mock function with given fields: _a0, _a1
func (_m *Repository) GetUserAuthenticateData(_a0 context.Context, _a1 string) (*model.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: _a0, _a1
func (_m *Repository) RegisterUser(_a0 context.Context, _a1 *request.RegisterRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.RegisterRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
