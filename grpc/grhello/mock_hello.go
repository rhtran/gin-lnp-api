// Code generated by MockGen. DO NOT EDIT.
// Source: grhello/hello.pb.go

// Package grhello is a generated GoMock package.
package grhello

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockHelloServiceClient is a mock of HelloServiceClient interface
type MockHelloServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockHelloServiceClientMockRecorder
}

// MockHelloServiceClientMockRecorder is the mock recorder for MockHelloServiceClient
type MockHelloServiceClientMockRecorder struct {
	mock *MockHelloServiceClient
}

// NewMockHelloServiceClient creates a new mock instance
func NewMockHelloServiceClient(ctrl *gomock.Controller) *MockHelloServiceClient {
	mock := &MockHelloServiceClient{ctrl: ctrl}
	mock.recorder = &MockHelloServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHelloServiceClient) EXPECT() *MockHelloServiceClientMockRecorder {
	return m.recorder
}

// Greet mocks base method
func (m *MockHelloServiceClient) Greet(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Greet", varargs...)
	ret0, _ := ret[0].(*Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Greet indicates an expected call of Greet
func (mr *MockHelloServiceClientMockRecorder) Greet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Greet", reflect.TypeOf((*MockHelloServiceClient)(nil).Greet), varargs...)
}

// MockHelloServiceServer is a mock of HelloServiceServer interface
type MockHelloServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockHelloServiceServerMockRecorder
}

// MockHelloServiceServerMockRecorder is the mock recorder for MockHelloServiceServer
type MockHelloServiceServerMockRecorder struct {
	mock *MockHelloServiceServer
}

// NewMockHelloServiceServer creates a new mock instance
func NewMockHelloServiceServer(ctrl *gomock.Controller) *MockHelloServiceServer {
	mock := &MockHelloServiceServer{ctrl: ctrl}
	mock.recorder = &MockHelloServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHelloServiceServer) EXPECT() *MockHelloServiceServerMockRecorder {
	return m.recorder
}

// Greet mocks base method
func (m *MockHelloServiceServer) Greet(arg0 context.Context, arg1 *Request) (*Response, error) {
	ret := m.ctrl.Call(m, "Greet", arg0, arg1)
	ret0, _ := ret[0].(*Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Greet indicates an expected call of Greet
func (mr *MockHelloServiceServerMockRecorder) Greet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Greet", reflect.TypeOf((*MockHelloServiceServer)(nil).Greet), arg0, arg1)
}
