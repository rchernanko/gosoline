// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import crud "github.com/applike/gosoline/pkg/apiserver/crud"
import db_repo "github.com/applike/gosoline/pkg/db-repo"
import mock "github.com/stretchr/testify/mock"

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// GetCreateInput provides a mock function with given fields:
func (_m *Handler) GetCreateInput() interface{} {
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

// GetModel provides a mock function with given fields:
func (_m *Handler) GetModel() db_repo.ModelBased {
	ret := _m.Called()

	var r0 db_repo.ModelBased
	if rf, ok := ret.Get(0).(func() db_repo.ModelBased); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db_repo.ModelBased)
		}
	}

	return r0
}

// GetRepository provides a mock function with given fields:
func (_m *Handler) GetRepository() crud.Repository {
	ret := _m.Called()

	var r0 crud.Repository
	if rf, ok := ret.Get(0).(func() crud.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crud.Repository)
		}
	}

	return r0
}

// GetUpdateInput provides a mock function with given fields:
func (_m *Handler) GetUpdateInput() interface{} {
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

// List provides a mock function with given fields: ctx, qb, apiView
func (_m *Handler) List(ctx context.Context, qb *db_repo.QueryBuilder, apiView string) (interface{}, error) {
	ret := _m.Called(ctx, qb, apiView)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, string) interface{}); ok {
		r0 = rf(ctx, qb, apiView)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *db_repo.QueryBuilder, string) error); ok {
		r1 = rf(ctx, qb, apiView)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransformCreate provides a mock function with given fields: input, model
func (_m *Handler) TransformCreate(input interface{}, model db_repo.ModelBased) error {
	ret := _m.Called(input, model)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, db_repo.ModelBased) error); ok {
		r0 = rf(input, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransformOutput provides a mock function with given fields: model, apiView
func (_m *Handler) TransformOutput(model db_repo.ModelBased, apiView string) (interface{}, error) {
	ret := _m.Called(model, apiView)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(db_repo.ModelBased, string) interface{}); ok {
		r0 = rf(model, apiView)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db_repo.ModelBased, string) error); ok {
		r1 = rf(model, apiView)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransformUpdate provides a mock function with given fields: input, model
func (_m *Handler) TransformUpdate(input interface{}, model db_repo.ModelBased) error {
	ret := _m.Called(input, model)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, db_repo.ModelBased) error); ok {
		r0 = rf(input, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
