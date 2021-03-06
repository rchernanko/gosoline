// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apiserver "github.com/applike/gosoline/pkg/apiserver"
import context "context"
import mock "github.com/stretchr/testify/mock"

// HandlerWithInput is an autogenerated mock type for the HandlerWithInput type
type HandlerWithInput struct {
	mock.Mock
}

// GetInput provides a mock function with given fields:
func (_m *HandlerWithInput) GetInput() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Handle provides a mock function with given fields: requestContext, request
func (_m *HandlerWithInput) Handle(requestContext context.Context, request *apiserver.Request) (*apiserver.Response, error) {
	ret := _m.Called(requestContext, request)

	var r0 *apiserver.Response
	if rf, ok := ret.Get(0).(func(context.Context, *apiserver.Request) *apiserver.Response); ok {
		r0 = rf(requestContext, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiserver.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *apiserver.Request) error); ok {
		r1 = rf(requestContext, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
