// Code generated by mockery v2.45.1. DO NOT EDIT.

package domainmocks

import mock "github.com/stretchr/testify/mock"

// EmailService is an autogenerated mock type for the EmailService type
type EmailService struct {
	mock.Mock
}

// SendEmail provides a mock function with given fields: to, subject, body
func (_m *EmailService) SendEmail(to string, subject string, body string) error {
	ret := _m.Called(to, subject, body)

	if len(ret) == 0 {
		panic("no return value specified for SendEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(to, subject, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendPasswordResetEmail provides a mock function with given fields: to, token
func (_m *EmailService) SendPasswordResetEmail(to string, token string) error {
	ret := _m.Called(to, token)

	if len(ret) == 0 {
		panic("no return value specified for SendPasswordResetEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(to, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEmailService creates a new instance of EmailService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmailService(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmailService {
	mock := &EmailService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
