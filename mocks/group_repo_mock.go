// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_groups is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "main/internal/models"
	reflect "reflect"
)

// MockGroupRepository is a mock of GroupRepository interface
type MockGroupRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGroupRepositoryMockRecorder
}

// MockGroupRepositoryMockRecorder is the mock recorder for MockGroupRepository
type MockGroupRepositoryMockRecorder struct {
	mock *MockGroupRepository
}

// NewMockGroupRepository creates a new mock instance
func NewMockGroupRepository(ctrl *gomock.Controller) *MockGroupRepository {
	mock := &MockGroupRepository{ctrl: ctrl}
	mock.recorder = &MockGroupRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGroupRepository) EXPECT() *MockGroupRepositoryMockRecorder {
	return m.recorder
}

// JoinTheGroup mocks base method
func (m *MockGroupRepository) JoinTheGroup(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JoinTheGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// JoinTheGroup indicates an expected call of JoinTheGroup
func (mr *MockGroupRepositoryMockRecorder) JoinTheGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JoinTheGroup", reflect.TypeOf((*MockGroupRepository)(nil).JoinTheGroup), arg0, arg1)
}

// CreateGroup mocks base method
func (m *MockGroupRepository) CreateGroup(arg0 int, arg1 models.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGroup indicates an expected call of CreateGroup
func (mr *MockGroupRepositoryMockRecorder) CreateGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockGroupRepository)(nil).CreateGroup), arg0, arg1)
}

// LeaveTheGroup mocks base method
func (m *MockGroupRepository) LeaveTheGroup(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveTheGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeaveTheGroup indicates an expected call of LeaveTheGroup
func (mr *MockGroupRepositoryMockRecorder) LeaveTheGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveTheGroup", reflect.TypeOf((*MockGroupRepository)(nil).LeaveTheGroup), arg0, arg1)
}

// IsUserOwnerOfGroup mocks base method
func (m *MockGroupRepository) IsUserOwnerOfGroup(arg0, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserOwnerOfGroup", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserOwnerOfGroup indicates an expected call of IsUserOwnerOfGroup
func (mr *MockGroupRepositoryMockRecorder) IsUserOwnerOfGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserOwnerOfGroup", reflect.TypeOf((*MockGroupRepository)(nil).IsUserOwnerOfGroup), arg0, arg1)
}

// CreatePostInGroup mocks base method
func (m *MockGroupRepository) CreatePostInGroup(arg0, arg1 int, arg2 models.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostInGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePostInGroup indicates an expected call of CreatePostInGroup
func (mr *MockGroupRepositoryMockRecorder) CreatePostInGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostInGroup", reflect.TypeOf((*MockGroupRepository)(nil).CreatePostInGroup), arg0, arg1, arg2)
}

// GetGroupProfile mocks base method
func (m *MockGroupRepository) GetGroupProfile(arg0, arg1 int) (models.GroupProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupProfile", arg0, arg1)
	ret0, _ := ret[0].(models.GroupProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupProfile indicates an expected call of GetGroupProfile
func (mr *MockGroupRepositoryMockRecorder) GetGroupProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupProfile", reflect.TypeOf((*MockGroupRepository)(nil).GetGroupProfile), arg0, arg1)
}

// GetGroupMembers mocks base method
func (m *MockGroupRepository) GetGroupMembers(arg0 int) ([]models.FriendLandingInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMembers", arg0)
	ret0, _ := ret[0].([]models.FriendLandingInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMembers indicates an expected call of GetGroupMembers
func (mr *MockGroupRepositoryMockRecorder) GetGroupMembers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMembers", reflect.TypeOf((*MockGroupRepository)(nil).GetGroupMembers), arg0)
}

// GetGroupFeeds mocks base method
func (m *MockGroupRepository) GetGroupFeeds(arg0, arg1 int) ([]models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupFeeds", arg0, arg1)
	ret0, _ := ret[0].([]models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupFeeds indicates an expected call of GetGroupFeeds
func (mr *MockGroupRepositoryMockRecorder) GetGroupFeeds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupFeeds", reflect.TypeOf((*MockGroupRepository)(nil).GetGroupFeeds), arg0, arg1)
}

// GetUserGroupsList mocks base method
func (m *MockGroupRepository) GetUserGroupsList(arg0 int) ([]models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserGroupsList", arg0)
	ret0, _ := ret[0].([]models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserGroupsList indicates an expected call of GetUserGroupsList
func (mr *MockGroupRepositoryMockRecorder) GetUserGroupsList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserGroupsList", reflect.TypeOf((*MockGroupRepository)(nil).GetUserGroupsList), arg0)
}

// SearchAllGroups mocks base method
func (m *MockGroupRepository) SearchAllGroups(arg0 int, arg1 string) ([]models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAllGroups", arg0, arg1)
	ret0, _ := ret[0].([]models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAllGroups indicates an expected call of SearchAllGroups
func (mr *MockGroupRepositoryMockRecorder) SearchAllGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAllGroups", reflect.TypeOf((*MockGroupRepository)(nil).SearchAllGroups), arg0, arg1)
}

// UpdateGroupProfile mocks base method
func (m *MockGroupRepository) UpdateGroupProfile(arg0, arg1 int, arg2 models.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroupProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGroupProfile indicates an expected call of UpdateGroupProfile
func (mr *MockGroupRepositoryMockRecorder) UpdateGroupProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupProfile", reflect.TypeOf((*MockGroupRepository)(nil).UpdateGroupProfile), arg0, arg1, arg2)
}
