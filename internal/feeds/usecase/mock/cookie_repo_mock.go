// Code generated by MockGen. DO NOT EDIT.
// Source: ../cookies/repository.go

// Package mock_cookies is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockCookieRepository is a mock of CookieRepository interface
type MockCookieRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCookieRepositoryMockRecorder
}

// MockCookieRepositoryMockRecorder is the mock recorder for MockCookieRepository
type MockCookieRepositoryMockRecorder struct {
	mock *MockCookieRepository
}

// NewMockCookieRepository creates a new mock instance
func NewMockCookieRepository(ctrl *gomock.Controller) *MockCookieRepository {
	mock := &MockCookieRepository{ctrl: ctrl}
	mock.recorder = &MockCookieRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCookieRepository) EXPECT() *MockCookieRepositoryMockRecorder {
	return m.recorder
}

// AddCookie mocks base method
func (m *MockCookieRepository) AddCookie(arg0 int, arg1 string, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCookie", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCookie indicates an expected call of AddCookie
func (mr *MockCookieRepositoryMockRecorder) AddCookie(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCookie", reflect.TypeOf((*MockCookieRepository)(nil).AddCookie), arg0, arg1, arg2)
}

// DeleteCookie mocks base method
func (m *MockCookieRepository) DeleteCookie(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCookie", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCookie indicates an expected call of DeleteCookie
func (mr *MockCookieRepositoryMockRecorder) DeleteCookie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCookie", reflect.TypeOf((*MockCookieRepository)(nil).DeleteCookie), arg0)
}

// GetUserIdByCookie mocks base method
func (m *MockCookieRepository) GetUserIdByCookie(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIdByCookie", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIdByCookie indicates an expected call of GetUserIdByCookie
func (mr *MockCookieRepositoryMockRecorder) GetUserIdByCookie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIdByCookie", reflect.TypeOf((*MockCookieRepository)(nil).GetUserIdByCookie), arg0)
}
