// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domaincategory "svc-receipt-luscious/core/domain/category"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteCategory provides a mock function with given fields: ctx, categoryID
func (_m *Repository) DeleteCategory(ctx context.Context, categoryID string) error {
	ret := _m.Called(ctx, categoryID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, categoryID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllListCategory provides a mock function with given fields: categoryName
func (_m *Repository) GetAllListCategory(categoryName string) ([]domaincategory.Category, error) {
	ret := _m.Called(categoryName)

	var r0 []domaincategory.Category
	if rf, ok := ret.Get(0).(func(string) []domaincategory.Category); ok {
		r0 = rf(categoryName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domaincategory.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(categoryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailCategory provides a mock function with given fields: categoryID
func (_m *Repository) GetDetailCategory(categoryID string) (*domaincategory.Category, error) {
	ret := _m.Called(categoryID)

	var r0 *domaincategory.Category
	if rf, ok := ret.Get(0).(func(string) *domaincategory.Category); ok {
		r0 = rf(categoryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domaincategory.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(categoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertCategory provides a mock function with given fields: ctx, inData
func (_m *Repository) InsertCategory(ctx context.Context, inData *domaincategory.Category) error {
	ret := _m.Called(ctx, inData)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domaincategory.Category) error); ok {
		r0 = rf(ctx, inData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCategory provides a mock function with given fields: ctx, inData
func (_m *Repository) UpdateCategory(ctx context.Context, inData *domaincategory.Category) error {
	ret := _m.Called(ctx, inData)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domaincategory.Category) error); ok {
		r0 = rf(ctx, inData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}