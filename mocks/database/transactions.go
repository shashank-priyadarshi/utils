// Code generated by mockery v2.42.1. DO NOT EDIT.

package ports

import (
	context "context"

	models "go.ssnk.in/utils/database/models"
	mock "github.com/stretchr/testify/mock"
)

// MockTransactions is an autogenerated mock type for the Transactions type
type MockTransactions struct {
	mock.Mock
}

type MockTransactions_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransactions) EXPECT() *MockTransactions_Expecter {
	return &MockTransactions_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields: _a0, _a1
func (_m *MockTransactions) Begin(_a0 context.Context, _a1 ...interface{}) (*models.Response, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 *models.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) (*models.Response, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) *models.Response); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...interface{}) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransactions_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type MockTransactions_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...interface{}
func (_e *MockTransactions_Expecter) Begin(_a0 interface{}, _a1 ...interface{}) *MockTransactions_Begin_Call {
	return &MockTransactions_Begin_Call{Call: _e.mock.On("Begin",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *MockTransactions_Begin_Call) Run(run func(_a0 context.Context, _a1 ...interface{})) *MockTransactions_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockTransactions_Begin_Call) Return(_a0 *models.Response, _a1 error) *MockTransactions_Begin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransactions_Begin_Call) RunAndReturn(run func(context.Context, ...interface{}) (*models.Response, error)) *MockTransactions_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MockTransactions) Execute(_a0 context.Context, _a1 ...interface{}) (*models.Response, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *models.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) (*models.Response, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) *models.Response); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...interface{}) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransactions_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockTransactions_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...interface{}
func (_e *MockTransactions_Expecter) Execute(_a0 interface{}, _a1 ...interface{}) *MockTransactions_Execute_Call {
	return &MockTransactions_Execute_Call{Call: _e.mock.On("Execute",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *MockTransactions_Execute_Call) Run(run func(_a0 context.Context, _a1 ...interface{})) *MockTransactions_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockTransactions_Execute_Call) Return(_a0 *models.Response, _a1 error) *MockTransactions_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransactions_Execute_Call) RunAndReturn(run func(context.Context, ...interface{}) (*models.Response, error)) *MockTransactions_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with given fields: _a0, _a1
func (_m *MockTransactions) Rollback(_a0 context.Context, _a1 ...interface{}) (*models.Response, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 *models.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) (*models.Response, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) *models.Response); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...interface{}) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransactions_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type MockTransactions_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...interface{}
func (_e *MockTransactions_Expecter) Rollback(_a0 interface{}, _a1 ...interface{}) *MockTransactions_Rollback_Call {
	return &MockTransactions_Rollback_Call{Call: _e.mock.On("Rollback",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *MockTransactions_Rollback_Call) Run(run func(_a0 context.Context, _a1 ...interface{})) *MockTransactions_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockTransactions_Rollback_Call) Return(_a0 *models.Response, _a1 error) *MockTransactions_Rollback_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransactions_Rollback_Call) RunAndReturn(run func(context.Context, ...interface{}) (*models.Response, error)) *MockTransactions_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTransactions creates a new instance of MockTransactions. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransactions(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransactions {
	mock := &MockTransactions{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
