// Code generated by MockGen. DO NOT EDIT.
// Source: random/random.go

// Package random is a generated GoMock package.
package random

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIFRandom is a mock of IFRandom interface.
type MockIFRandom struct {
	ctrl     *gomock.Controller
	recorder *MockIFRandomMockRecorder
}

// MockIFRandomMockRecorder is the mock recorder for MockIFRandom.
type MockIFRandomMockRecorder struct {
	mock *MockIFRandom
}

// NewMockIFRandom creates a new mock instance.
func NewMockIFRandom(ctrl *gomock.Controller) *MockIFRandom {
	mock := &MockIFRandom{ctrl: ctrl}
	mock.recorder = &MockIFRandomMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFRandom) EXPECT() *MockIFRandomMockRecorder {
	return m.recorder
}

// Intn mocks base method.
func (m *MockIFRandom) Intn(n int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intn", n)
	ret0, _ := ret[0].(int)
	return ret0
}

// Intn indicates an expected call of Intn.
func (mr *MockIFRandomMockRecorder) Intn(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intn", reflect.TypeOf((*MockIFRandom)(nil).Intn), n)
}
