// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"

	model "github.com/scys12/clean-architecture-golang/model"

	redis "github.com/gomodule/redigo/redis"

	session "github.com/scys12/clean-architecture-golang/pkg/session"
)

// SessionStore is an autogenerated mock type for the SessionStore type
type SessionStore struct {
	mock.Mock
}

// Connect provides a mock function with given fields:
func (_m *SessionStore) Connect() redis.Conn {
	ret := _m.Called()

	var r0 redis.Conn
	if rf, ok := ret.Get(0).(func() redis.Conn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(redis.Conn)
		}
	}

	return r0
}

// CreateSession provides a mock function with given fields: _a0, _a1
func (_m *SessionStore) CreateSession(_a0 echo.Context, _a1 *model.User) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, *model.User) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *SessionStore) Get(_a0 string) (session.Session, error) {
	ret := _m.Called(_a0)

	var r0 session.Session
	if rf, ok := ret.Get(0).(func(string) session.Session); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(session.Session)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: _a0, _a1
func (_m *SessionStore) Set(_a0 string, _a1 session.Session) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, session.Session) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
