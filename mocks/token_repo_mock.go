// Code generated by MockGen. DO NOT EDIT.
// Source: main/internal/socket_token (interfaces: TokenRepository)

// Package mock_socket_token is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTokenRepository is a mock of TokenRepository interface
type MockTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTokenRepositoryMockRecorder
}

// MockTokenRepositoryMockRecorder is the mock recorder for MockTokenRepository
type MockTokenRepositoryMockRecorder struct {
	mock *MockTokenRepository
}

// NewMockTokenRepository creates a new mock instance
func NewMockTokenRepository(ctrl *gomock.Controller) *MockTokenRepository {
	mock := &MockTokenRepository{ctrl: ctrl}
	mock.recorder = &MockTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenRepository) EXPECT() *MockTokenRepositoryMockRecorder {
	return m.recorder
}

// AddNewToken mocks base method
func (m *MockTokenRepository) AddNewToken(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewToken indicates an expected call of AddNewToken
func (mr *MockTokenRepositoryMockRecorder) AddNewToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewToken", reflect.TypeOf((*MockTokenRepository)(nil).AddNewToken), arg0, arg1)
}

// GetUserIdByToken mocks base method
func (m *MockTokenRepository) GetUserIdByToken(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIdByToken", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIdByToken indicates an expected call of GetUserIdByToken
func (mr *MockTokenRepositoryMockRecorder) GetUserIdByToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIdByToken", reflect.TypeOf((*MockTokenRepository)(nil).GetUserIdByToken), arg0)
}
