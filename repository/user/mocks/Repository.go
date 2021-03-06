// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/scys12/clean-architecture-golang/model"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// EditUserProfile provides a mock function with given fields: _a0, _a1
func (_m *Repository) EditUserProfile(_a0 context.Context, _a1 model.UserProfile) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UserProfile) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserAuthenticateData provides a mock function with given fields: _a0, _a1
func (_m *Repository) GetUserAuthenticateData(_a0 context.Context, _a1 map[string]interface{}) (*model.UserAuth, *model.UserProfile, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.UserAuth
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) *model.UserAuth); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserAuth)
		}
	}

	var r1 *model.UserProfile
	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) *model.UserProfile); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.UserProfile)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, map[string]interface{}) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RegisterUser provides a mock function with given fields: _a0, _a1
func (_m *Repository) RegisterUser(_a0 context.Context, _a1 model.UserAuth) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UserAuth) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
