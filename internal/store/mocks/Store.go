// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "github.com/supasiti/prac-aws-serverless-go/internal/store/user"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

type Store_Expecter struct {
	mock *mock.Mock
}

func (_m *Store) EXPECT() *Store_Expecter {
	return &Store_Expecter{mock: &_m.Mock}
}

// GetUser provides a mock function with given fields: _a0, _a1
func (_m *Store) GetUser(_a0 context.Context, _a1 int) (*user.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*user.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *user.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type Store_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int
func (_e *Store_Expecter) GetUser(_a0 interface{}, _a1 interface{}) *Store_GetUser_Call {
	return &Store_GetUser_Call{Call: _e.mock.On("GetUser", _a0, _a1)}
}

func (_c *Store_GetUser_Call) Run(run func(_a0 context.Context, _a1 int)) *Store_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Store_GetUser_Call) Return(_a0 *user.User, _a1 error) *Store_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_GetUser_Call) RunAndReturn(run func(context.Context, int) (*user.User, error)) *Store_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStore(t mockConstructorTestingTNewStore) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}