// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-flotta/flotta-operator/internal/k8sclient (interfaces: K8sClient)

// Package k8sclient is a generated GoMock package.
package k8sclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "k8s.io/apimachinery/pkg/types"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockK8sClient is a mock of K8sClient interface.
type MockK8sClient struct {
	ctrl     *gomock.Controller
	recorder *MockK8sClientMockRecorder
}

// MockK8sClientMockRecorder is the mock recorder for MockK8sClient.
type MockK8sClientMockRecorder struct {
	mock *MockK8sClient
}

// NewMockK8sClient creates a new mock instance.
func NewMockK8sClient(ctrl *gomock.Controller) *MockK8sClient {
	mock := &MockK8sClient{ctrl: ctrl}
	mock.recorder = &MockK8sClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8sClient) EXPECT() *MockK8sClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockK8sClient) Get(arg0 context.Context, arg1 types.NamespacedName, arg2 client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockK8sClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockK8sClient)(nil).Get), arg0, arg1, arg2)
}
