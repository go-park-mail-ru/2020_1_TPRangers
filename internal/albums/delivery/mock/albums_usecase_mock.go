package mock

import (
	gomock "github.com/golang/mock/gomock"
)

// MockFriendUseCase is a mock of FriendUseCase interface
type MockAlbumUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumUseCaseMockRecorder
}

// MockFriendUseCaseMockRecorder is the mock recorder for MockFriendUseCase
type MockAlbumUseCaseMockRecorder struct {
	mock *MockAlbumUseCase
}

// NewMockFriendUseCase creates a new mock instance
func NewMockAlbumUseCase(ctrl *gomock.Controller) *MockAlbumUseCase {
	mock := &MockAlbumUseCase{ctrl: ctrl}
	mock.recorder = &MockAlbumUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlbumUseCase) EXPECT() *MockAlbumUseCaseMockRecorder {
	return m.recorder
}

// AddFriend mocks base method
func (m *MockAlbumUseCase) AddFriend(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFriend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}