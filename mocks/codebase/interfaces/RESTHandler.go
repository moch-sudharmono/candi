// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo"

	mock "github.com/stretchr/testify/mock"
)

// RESTHandler is an autogenerated mock type for the RESTHandler type
type RESTHandler struct {
	mock.Mock
}

// Mount provides a mock function with given fields: group
func (_m *RESTHandler) Mount(group *echo.Group) {
	_m.Called(group)
}
