// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/Optum/dce/pkg/model"

// WriterDeleter is an autogenerated mock type for the WriterDeleter type
type WriterDeleter struct {
	mock.Mock
}

// DeleteAccount provides a mock function with given fields: input
func (_m *WriterDeleter) DeleteAccount(input *model.Account) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Account) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteAccount provides a mock function with given fields: input, lastModifiedOn
func (_m *WriterDeleter) WriteAccount(input *model.Account, lastModifiedOn *int64) error {
	ret := _m.Called(input, lastModifiedOn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Account, *int64) error); ok {
		r0 = rf(input, lastModifiedOn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
