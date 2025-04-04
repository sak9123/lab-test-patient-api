// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	errs "hospitalApi/pkg/errs"

	mock "github.com/stretchr/testify/mock"

	model "hospitalApi/pkg/model"
)

// IStaffService is an autogenerated mock type for the IStaffService type
type IStaffService struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *IStaffService) Create(input model.Staff) (bool, *errs.Error) {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 bool
	var r1 *errs.Error
	if rf, ok := ret.Get(0).(func(model.Staff) (bool, *errs.Error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(model.Staff) bool); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(model.Staff) *errs.Error); ok {
		r1 = rf(input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.Error)
		}
	}

	return r0, r1
}

// IsExistsUsername provides a mock function with given fields: username, hospitalCode
func (_m *IStaffService) IsExistsUsername(username string, hospitalCode string) *errs.Error {
	ret := _m.Called(username, hospitalCode)

	if len(ret) == 0 {
		panic("no return value specified for IsExistsUsername")
	}

	var r0 *errs.Error
	if rf, ok := ret.Get(0).(func(string, string) *errs.Error); ok {
		r0 = rf(username, hospitalCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.Error)
		}
	}

	return r0
}

// Login provides a mock function with given fields: input
func (_m *IStaffService) Login(input model.StaffCriteria) (string, *errs.Error) {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 *errs.Error
	if rf, ok := ret.Get(0).(func(model.StaffCriteria) (string, *errs.Error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(model.StaffCriteria) string); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(model.StaffCriteria) *errs.Error); ok {
		r1 = rf(input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.Error)
		}
	}

	return r0, r1
}

// ValidateLogin provides a mock function with given fields: input
func (_m *IStaffService) ValidateLogin(input model.StaffCriteria) *errs.Error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for ValidateLogin")
	}

	var r0 *errs.Error
	if rf, ok := ret.Get(0).(func(model.StaffCriteria) *errs.Error); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.Error)
		}
	}

	return r0
}

// ValidateSave provides a mock function with given fields: input
func (_m *IStaffService) ValidateSave(input model.Staff) *errs.Error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for ValidateSave")
	}

	var r0 *errs.Error
	if rf, ok := ret.Get(0).(func(model.Staff) *errs.Error); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.Error)
		}
	}

	return r0
}

// NewIStaffService creates a new instance of IStaffService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIStaffService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IStaffService {
	mock := &IStaffService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
