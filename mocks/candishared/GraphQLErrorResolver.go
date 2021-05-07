// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GraphQLErrorResolver is an autogenerated mock type for the GraphQLErrorResolver type
type GraphQLErrorResolver struct {
	mock.Mock
}

// Error provides a mock function with given fields:
func (_m *GraphQLErrorResolver) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Extensions provides a mock function with given fields:
func (_m *GraphQLErrorResolver) Extensions() map[string]interface{} {
	ret := _m.Called()

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}
