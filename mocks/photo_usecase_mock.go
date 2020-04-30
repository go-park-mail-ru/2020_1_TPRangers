// Code generated by MockGen. DO NOT EDIT.

// Source: main/internal/photos (interfaces: PhotoUseCase)

// Package mock_photos is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "main/models"
	reflect "reflect"
)

// MockPhotoUseCase is a mock of PhotoUseCase interface
type MockPhotoUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoUseCaseMockRecorder
}

// MockPhotoUseCaseMockRecorder is the mock recorder for MockPhotoUseCase
type MockPhotoUseCaseMockRecorder struct {
	mock *MockPhotoUseCase
}

// NewMockPhotoUseCase creates a new mock instance
func NewMockPhotoUseCase(ctrl *gomock.Controller) *MockPhotoUseCase {
	mock := &MockPhotoUseCase{ctrl: ctrl}
	mock.recorder = &MockPhotoUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPhotoUseCase) EXPECT() *MockPhotoUseCaseMockRecorder {
	return m.recorder
}

// GetPhotosFromAlbum mocks base method
func (m *MockPhotoUseCase) GetPhotosFromAlbum(arg0 int) (models.Photos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotosFromAlbum", arg0)
	ret0, _ := ret[0].(models.Photos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotosFromAlbum indicates an expected call of GetPhotosFromAlbum
func (mr *MockPhotoUseCaseMockRecorder) GetPhotosFromAlbum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotosFromAlbum", reflect.TypeOf((*MockPhotoUseCase)(nil).GetPhotosFromAlbum), arg0)
}

// UploadPhotoToAlbum mocks base method
func (m *MockPhotoUseCase) UploadPhotoToAlbum(arg0 models.PhotoInAlbum) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadPhotoToAlbum", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadPhotoToAlbum indicates an expected call of UploadPhotoToAlbum
func (mr *MockPhotoUseCaseMockRecorder) UploadPhotoToAlbum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPhotoToAlbum", reflect.TypeOf((*MockPhotoUseCase)(nil).UploadPhotoToAlbum), arg0)
}
