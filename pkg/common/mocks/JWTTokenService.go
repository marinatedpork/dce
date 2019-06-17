// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import jwt "github.com/dgrijalva/jwt-go"
import mock "github.com/stretchr/testify/mock"

// JWTTokenService is an autogenerated mock type for the JWTTokenService type
type JWTTokenService struct {
	mock.Mock
}

// ParseJWT provides a mock function with given fields:
func (_m *JWTTokenService) ParseJWT() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// getKey provides a mock function with given fields: token
func (_m *JWTTokenService) getKey(token *jwt.Token) (interface{}, error) {
	ret := _m.Called(token)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*jwt.Token) interface{}); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*jwt.Token) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getTID provides a mock function with given fields: token
func (_m *JWTTokenService) getTID(token string) (string, error) {
	ret := _m.Called(token)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}