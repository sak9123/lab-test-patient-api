// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	entity "hospitalApi/pkg/entity"

	mock "github.com/stretchr/testify/mock"

	model "hospitalApi/pkg/model"
)

// IStaffsRepository is an autogenerated mock type for the IStaffsRepository type
type IStaffsRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *IStaffsRepository) Create(input entity.Staff) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Staff) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: input
func (_m *IStaffsRepository) Get(input model.StaffCriteria) ([]entity.Staff, error) {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []entity.Staff
	var r1 error
	if rf, ok := ret.Get(0).(func(model.StaffCriteria) ([]entity.Staff, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(model.StaffCriteria) []entity.Staff); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Staff)
		}
	}

	if rf, ok := ret.Get(1).(func(model.StaffCriteria) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIStaffsRepository creates a new instance of IStaffsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIStaffsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IStaffsRepository {
	mock := &IStaffsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
