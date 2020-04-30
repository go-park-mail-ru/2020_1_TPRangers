// Code generated by MockGen. DO NOT EDIT.
// Source: ../repository.go

// Package mock_albums is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "main/internal/models"
	reflect "reflect"
)

// MockAlbumRepository is a mock of AlbumRepository interface
type MockAlbumRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumRepositoryMockRecorder
}

// MockAlbumRepositoryMockRecorder is the mock recorder for MockAlbumRepository
type MockAlbumRepositoryMockRecorder struct {
	mock *MockAlbumRepository
}

// NewMockAlbumRepository creates a new mock instance
func NewMockAlbumRepository(ctrl *gomock.Controller) *MockAlbumRepository {
	mock := &MockAlbumRepository{ctrl: ctrl}
	mock.recorder = &MockAlbumRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlbumRepository) EXPECT() *MockAlbumRepositoryMockRecorder {
	return m.recorder
}

// GetAlbums mocks base method
func (m *MockAlbumRepository) GetAlbums(arg0 int) ([]models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbums", arg0)
	ret0, _ := ret[0].([]models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbums indicates an expected call of GetAlbums
func (mr *MockAlbumRepositoryMockRecorder) GetAlbums(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbums", reflect.TypeOf((*MockAlbumRepository)(nil).GetAlbums), arg0)
}

// CreateAlbum mocks base method
func (m *MockAlbumRepository) CreateAlbum(arg0 int, arg1 models.AlbumReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlbum", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAlbum indicates an expected call of CreateAlbum
func (mr *MockAlbumRepositoryMockRecorder) CreateAlbum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlbum", reflect.TypeOf((*MockAlbumRepository)(nil).CreateAlbum), arg0, arg1)
}
