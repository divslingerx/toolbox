// Code generated by mockery v2.40.3. DO NOT EDIT.

package home

import (
	context "context"

	domain "toolbox/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// List provides a mock function with given fields: ctx
func (_m *MockService) List(ctx context.Context) ([]*domain.Todo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*domain.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.Todo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Todo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter) List(ctx interface{}) *MockService_List_Call {
	return &MockService_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockService_List_Call) Run(run func(ctx context.Context)) *MockService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_List_Call) Return(_a0 []*domain.Todo, _a1 error) *MockService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_List_Call) RunAndReturn(run func(context.Context) ([]*domain.Todo, error)) *MockService_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
