// Code generated by MockGen. DO NOT EDIT.
// Source: main/internal/microservices/likes/delivery (interfaces: LikeCheckerClient)

// Package mock_delivery is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	delivery "main/internal/microservices/likes/delivery"
	reflect "reflect"
)

// MockLikeCheckerClient is a mock of LikeCheckerClient interface
type MockLikeCheckerClient struct {
	ctrl     *gomock.Controller
	recorder *MockLikeCheckerClientMockRecorder
}

// MockLikeCheckerClientMockRecorder is the mock recorder for MockLikeCheckerClient
type MockLikeCheckerClientMockRecorder struct {
	mock *MockLikeCheckerClient
}

// NewMockLikeCheckerClient creates a new mock instance
func NewMockLikeCheckerClient(ctrl *gomock.Controller) *MockLikeCheckerClient {
	mock := &MockLikeCheckerClient{ctrl: ctrl}
	mock.recorder = &MockLikeCheckerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLikeCheckerClient) EXPECT() *MockLikeCheckerClientMockRecorder {
	return m.recorder
}

// DislikeComment mocks base method
func (m *MockLikeCheckerClient) DislikeComment(arg0 context.Context, arg1 *delivery.Like, arg2 ...grpc.CallOption) (*delivery.Dummy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DislikeComment", varargs...)
	ret0, _ := ret[0].(*delivery.Dummy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DislikeComment indicates an expected call of DislikeComment
func (mr *MockLikeCheckerClientMockRecorder) DislikeComment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DislikeComment", reflect.TypeOf((*MockLikeCheckerClient)(nil).DislikeComment), varargs...)
}

// DislikePhoto mocks base method
func (m *MockLikeCheckerClient) DislikePhoto(arg0 context.Context, arg1 *delivery.Like, arg2 ...grpc.CallOption) (*delivery.Dummy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DislikePhoto", varargs...)
	ret0, _ := ret[0].(*delivery.Dummy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DislikePhoto indicates an expected call of DislikePhoto
func (mr *MockLikeCheckerClientMockRecorder) DislikePhoto(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DislikePhoto", reflect.TypeOf((*MockLikeCheckerClient)(nil).DislikePhoto), varargs...)
}

// DislikePost mocks base method
func (m *MockLikeCheckerClient) DislikePost(arg0 context.Context, arg1 *delivery.Like, arg2 ...grpc.CallOption) (*delivery.Dummy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DislikePost", varargs...)
	ret0, _ := ret[0].(*delivery.Dummy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DislikePost indicates an expected call of DislikePost
func (mr *MockLikeCheckerClientMockRecorder) DislikePost(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DislikePost", reflect.TypeOf((*MockLikeCheckerClient)(nil).DislikePost), varargs...)
}

// LikeComment mocks base method
func (m *MockLikeCheckerClient) LikeComment(arg0 context.Context, arg1 *delivery.Like, arg2 ...grpc.CallOption) (*delivery.Dummy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LikeComment", varargs...)
	ret0, _ := ret[0].(*delivery.Dummy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LikeComment indicates an expected call of LikeComment
func (mr *MockLikeCheckerClientMockRecorder) LikeComment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikeComment", reflect.TypeOf((*MockLikeCheckerClient)(nil).LikeComment), varargs...)
}

// LikePhoto mocks base method
func (m *MockLikeCheckerClient) LikePhoto(arg0 context.Context, arg1 *delivery.Like, arg2 ...grpc.CallOption) (*delivery.Dummy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LikePhoto", varargs...)
	ret0, _ := ret[0].(*delivery.Dummy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LikePhoto indicates an expected call of LikePhoto
func (mr *MockLikeCheckerClientMockRecorder) LikePhoto(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikePhoto", reflect.TypeOf((*MockLikeCheckerClient)(nil).LikePhoto), varargs...)
}

// LikePost mocks base method
func (m *MockLikeCheckerClient) LikePost(arg0 context.Context, arg1 *delivery.Like, arg2 ...grpc.CallOption) (*delivery.Dummy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LikePost", varargs...)
	ret0, _ := ret[0].(*delivery.Dummy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LikePost indicates an expected call of LikePost
func (mr *MockLikeCheckerClientMockRecorder) LikePost(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikePost", reflect.TypeOf((*MockLikeCheckerClient)(nil).LikePost), varargs...)
}
