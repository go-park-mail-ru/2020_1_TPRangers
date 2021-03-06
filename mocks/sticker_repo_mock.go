// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_stickers is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "main/internal/models"
	reflect "reflect"
)

// MockStickerRepo is a mock of StickerRepo interface
type MockStickerRepo struct {
	ctrl     *gomock.Controller
	recorder *MockStickerRepoMockRecorder
}

// MockStickerRepoMockRecorder is the mock recorder for MockStickerRepo
type MockStickerRepoMockRecorder struct {
	mock *MockStickerRepo
}

// NewMockStickerRepo creates a new mock instance
func NewMockStickerRepo(ctrl *gomock.Controller) *MockStickerRepo {
	mock := &MockStickerRepo{ctrl: ctrl}
	mock.recorder = &MockStickerRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStickerRepo) EXPECT() *MockStickerRepoMockRecorder {
	return m.recorder
}

// UploadStickerPack mocks base method
func (m *MockStickerRepo) UploadStickerPack(arg0 int, arg1 models.StickerPack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadStickerPack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadStickerPack indicates an expected call of UploadStickerPack
func (mr *MockStickerRepoMockRecorder) UploadStickerPack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadStickerPack", reflect.TypeOf((*MockStickerRepo)(nil).UploadStickerPack), arg0, arg1)
}

// GetStickerPacks mocks base method
func (m *MockStickerRepo) GetStickerPacks(arg0 int) ([]models.StickerPack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStickerPacks", arg0)
	ret0, _ := ret[0].([]models.StickerPack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStickerPacks indicates an expected call of GetStickerPacks
func (mr *MockStickerRepoMockRecorder) GetStickerPacks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStickerPacks", reflect.TypeOf((*MockStickerRepo)(nil).GetStickerPacks), arg0)
}

// PurchaseStickerPack mocks base method
func (m *MockStickerRepo) PurchaseStickerPack(arg0 int, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurchaseStickerPack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurchaseStickerPack indicates an expected call of PurchaseStickerPack
func (mr *MockStickerRepoMockRecorder) PurchaseStickerPack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurchaseStickerPack", reflect.TypeOf((*MockStickerRepo)(nil).PurchaseStickerPack), arg0, arg1)
}
