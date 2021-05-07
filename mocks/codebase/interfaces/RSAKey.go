// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	rsa "crypto/rsa"

	mock "github.com/stretchr/testify/mock"
)

// RSAKey is an autogenerated mock type for the RSAKey type
type RSAKey struct {
	mock.Mock
}

// PrivateKey provides a mock function with given fields:
func (_m *RSAKey) PrivateKey() *rsa.PrivateKey {
	ret := _m.Called()

	var r0 *rsa.PrivateKey
	if rf, ok := ret.Get(0).(func() *rsa.PrivateKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rsa.PrivateKey)
		}
	}

	return r0
}

// PublicKey provides a mock function with given fields:
func (_m *RSAKey) PublicKey() *rsa.PublicKey {
	ret := _m.Called()

	var r0 *rsa.PublicKey
	if rf, ok := ret.Get(0).(func() *rsa.PublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rsa.PublicKey)
		}
	}

	return r0
}
