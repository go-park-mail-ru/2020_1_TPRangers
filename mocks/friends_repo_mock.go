// Code generated by MockGen. DO NOT EDIT.
// Source: main/internal/friends (interfaces: FriendRepository)

// Package mock_friends is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "main/internal/models"
	reflect "reflect"
)

// MockFriendRepository is a mock of FriendRepository interface
type MockFriendRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFriendRepositoryMockRecorder
}

// MockFriendRepositoryMockRecorder is the mock recorder for MockFriendRepository
type MockFriendRepositoryMockRecorder struct {
	mock *MockFriendRepository
}

// NewMockFriendRepository creates a new mock instance
func NewMockFriendRepository(ctrl *gomock.Controller) *MockFriendRepository {
	mock := &MockFriendRepository{ctrl: ctrl}
	mock.recorder = &MockFriendRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFriendRepository) EXPECT() *MockFriendRepositoryMockRecorder {
	return m.recorder
}

// AddFriend mocks base method
func (m *MockFriendRepository) AddFriend(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFriend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFriend indicates an expected call of AddFriend
func (mr *MockFriendRepositoryMockRecorder) AddFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFriend", reflect.TypeOf((*MockFriendRepository)(nil).AddFriend), arg0, arg1)
}

// CheckFriendship mocks base method
func (m *MockFriendRepository) CheckFriendship(arg0, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFriendship", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckFriendship indicates an expected call of CheckFriendship
func (mr *MockFriendRepositoryMockRecorder) CheckFriendship(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFriendship", reflect.TypeOf((*MockFriendRepository)(nil).CheckFriendship), arg0, arg1)
}

// DeleteFriend mocks base method
func (m *MockFriendRepository) DeleteFriend(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFriend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFriend indicates an expected call of DeleteFriend
func (mr *MockFriendRepositoryMockRecorder) DeleteFriend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFriend", reflect.TypeOf((*MockFriendRepository)(nil).DeleteFriend), arg0, arg1)
}

// GetAllFriendsByLogin mocks base method
func (m *MockFriendRepository) GetAllFriendsByLogin(arg0 string) ([]models.FriendLandingInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFriendsByLogin", arg0)
	ret0, _ := ret[0].([]models.FriendLandingInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFriendsByLogin indicates an expected call of GetAllFriendsByLogin
func (mr *MockFriendRepositoryMockRecorder) GetAllFriendsByLogin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFriendsByLogin", reflect.TypeOf((*MockFriendRepository)(nil).GetAllFriendsByLogin), arg0)
}

// GetFriendIdByLogin mocks base method
func (m *MockFriendRepository) GetFriendIdByLogin(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriendIdByLogin", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriendIdByLogin indicates an expected call of GetFriendIdByLogin
func (mr *MockFriendRepositoryMockRecorder) GetFriendIdByLogin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriendIdByLogin", reflect.TypeOf((*MockFriendRepository)(nil).GetFriendIdByLogin), arg0)
}

// GetIdByLogin mocks base method
func (m *MockFriendRepository) GetIdByLogin(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdByLogin", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdByLogin indicates an expected call of GetIdByLogin
func (mr *MockFriendRepositoryMockRecorder) GetIdByLogin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdByLogin", reflect.TypeOf((*MockFriendRepository)(nil).GetIdByLogin), arg0)
}

// GetUserFriendsById mocks base method
func (m *MockFriendRepository) GetUserFriendsById(arg0, arg1 int) ([]models.FriendLandingInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFriendsById", arg0, arg1)
	ret0, _ := ret[0].([]models.FriendLandingInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFriendsById indicates an expected call of GetUserFriendsById
func (mr *MockFriendRepositoryMockRecorder) GetUserFriendsById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFriendsById", reflect.TypeOf((*MockFriendRepository)(nil).GetUserFriendsById), arg0, arg1)
}

// GetUserFriendsByLogin mocks base method
func (m *MockFriendRepository) GetUserFriendsByLogin(arg0 string, arg1 int) ([]models.FriendLandingInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFriendsByLogin", arg0, arg1)
	ret0, _ := ret[0].([]models.FriendLandingInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFriendsByLogin indicates an expected call of GetUserFriendsByLogin
func (mr *MockFriendRepositoryMockRecorder) GetUserFriendsByLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFriendsByLogin", reflect.TypeOf((*MockFriendRepository)(nil).GetUserFriendsByLogin), arg0, arg1)
}

// GetUserLoginById mocks base method
func (m *MockFriendRepository) GetUserLoginById(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLoginById", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLoginById indicates an expected call of GetUserLoginById
func (mr *MockFriendRepositoryMockRecorder) GetUserLoginById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLoginById", reflect.TypeOf((*MockFriendRepository)(nil).GetUserLoginById), arg0)
}

// SearchFriends mocks base method
func (m *MockFriendRepository) SearchFriends(arg0 int, arg1 string) ([]models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFriends", arg0, arg1)
	ret0, _ := ret[0].([]models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchFriends indicates an expected call of SearchFriends
func (mr *MockFriendRepositoryMockRecorder) SearchFriends(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFriends", reflect.TypeOf((*MockFriendRepository)(nil).SearchFriends), arg0, arg1)
}
