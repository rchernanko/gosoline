// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gin "github.com/gin-gonic/gin"
import mock "github.com/stretchr/testify/mock"

// Authenticator is an autogenerated mock type for the Authenticator type
type Authenticator struct {
	mock.Mock
}

// IsValid provides a mock function with given fields: ginCtx
func (_m *Authenticator) IsValid(ginCtx *gin.Context) (bool, error) {
	ret := _m.Called(ginCtx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*gin.Context) bool); ok {
		r0 = rf(ginCtx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ginCtx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
