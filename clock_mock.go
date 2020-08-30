// Code generated by MockGen. DO NOT EDIT.
// Source: clock.go

// Package main is a generated GoMock package.
package main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockClocker is a mock of Clocker interface
type MockClocker struct {
	ctrl     *gomock.Controller
	recorder *MockClockerMockRecorder
}

// MockClockerMockRecorder is the mock recorder for MockClocker
type MockClockerMockRecorder struct {
	mock *MockClocker
}

// NewMockClocker creates a new mock instance
func NewMockClocker(ctrl *gomock.Controller) *MockClocker {
	mock := &MockClocker{ctrl: ctrl}
	mock.recorder = &MockClockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClocker) EXPECT() *MockClockerMockRecorder {
	return m.recorder
}

// Now mocks base method
func (m *MockClocker) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now
func (mr *MockClockerMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClocker)(nil).Now))
}
