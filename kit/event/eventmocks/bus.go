// Code generated by mockery v2.45.1. DO NOT EDIT.

package eventmocks

import (
	context "context"
	event "user-service/kit/event"

	mock "github.com/stretchr/testify/mock"
)

// Bus is an autogenerated mock type for the Bus type
type Bus struct {
	mock.Mock
}

// Publish provides a mock function with given fields: _a0, _a1
func (_m *Bus) Publish(_a0 context.Context, _a1 []event.Event) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []event.Event) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBus creates a new instance of Bus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBus(t interface {
	mock.TestingT
	Cleanup(func())
}) *Bus {
	mock := &Bus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
