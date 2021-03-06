// Code generated by MockGen. DO NOT EDIT.
// Source: main/internal/chats (interfaces: ChatUseCase)
// Package mock_chats is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "main/internal/models"
	reflect "reflect"
)

// MockChatUseCase is a mock of ChatUseCase interface
type MockChatUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockChatUseCaseMockRecorder
}

// MockChatUseCaseMockRecorder is the mock recorder for MockChatUseCase
type MockChatUseCaseMockRecorder struct {
	mock *MockChatUseCase
}

// NewMockChatUseCase creates a new mock instance
func NewMockChatUseCase(ctrl *gomock.Controller) *MockChatUseCase {
	mock := &MockChatUseCase{ctrl: ctrl}
	mock.recorder = &MockChatUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatUseCase) EXPECT() *MockChatUseCaseMockRecorder {
	return m.recorder
}

// CreateChat mocks base method
func (m *MockChatUseCase) CreateChat(arg0 models.NewChatUsers, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChat", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateChat indicates an expected call of CreateChat
func (mr *MockChatUseCaseMockRecorder) CreateChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChat", reflect.TypeOf((*MockChatUseCase)(nil).CreateChat), arg0, arg1)
}

// ExitChat mocks base method
func (m *MockChatUseCase) ExitChat(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExitChat", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExitChat indicates an expected call of ExitChat
func (mr *MockChatUseCaseMockRecorder) ExitChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExitChat", reflect.TypeOf((*MockChatUseCase)(nil).ExitChat), arg0, arg1)
}

// GetAllChats mocks base method
func (m *MockChatUseCase) GetAllChats(arg0 int) ([]models.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChats", arg0)
	ret0, _ := ret[0].([]models.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllChats indicates an expected call of GetAllChats
func (mr *MockChatUseCaseMockRecorder) GetAllChats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChats", reflect.TypeOf((*MockChatUseCase)(nil).GetAllChats), arg0)
}

// GetChatMessages mocks base method
func (m *MockChatUseCase) GetChatMessages(arg0 string, arg1 int) (models.ChatAndMessages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatMessages", arg0, arg1)
	ret0, _ := ret[0].(models.ChatAndMessages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatMessages indicates an expected call of GetChatMessages
func (mr *MockChatUseCaseMockRecorder) GetChatMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatMessages", reflect.TypeOf((*MockChatUseCase)(nil).GetChatMessages), arg0, arg1)
}
