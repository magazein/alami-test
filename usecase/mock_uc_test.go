// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package usecase is a generated GoMock package.
package usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEndOfDayUCItf is a mock of EndOfDayUCItf interface.
type MockEndOfDayUCItf struct {
	ctrl     *gomock.Controller
	recorder *MockEndOfDayUCItfMockRecorder
}

// MockEndOfDayUCItfMockRecorder is the mock recorder for MockEndOfDayUCItf.
type MockEndOfDayUCItfMockRecorder struct {
	mock *MockEndOfDayUCItf
}

// NewMockEndOfDayUCItf creates a new mock instance.
func NewMockEndOfDayUCItf(ctrl *gomock.Controller) *MockEndOfDayUCItf {
	mock := &MockEndOfDayUCItf{ctrl: ctrl}
	mock.recorder = &MockEndOfDayUCItfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndOfDayUCItf) EXPECT() *MockEndOfDayUCItfMockRecorder {
	return m.recorder
}

// Proceed mocks base method.
func (m *MockEndOfDayUCItf) Proceed() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Proceed")
	ret0, _ := ret[0].(error)
	return ret0
}

// Proceed indicates an expected call of Proceed.
func (mr *MockEndOfDayUCItfMockRecorder) Proceed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proceed", reflect.TypeOf((*MockEndOfDayUCItf)(nil).Proceed))
}

// UpdateAvgBalance mocks base method.
func (m *MockEndOfDayUCItf) UpdateAvgBalance(i int, data []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvgBalance", i, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvgBalance indicates an expected call of UpdateAvgBalance.
func (mr *MockEndOfDayUCItfMockRecorder) UpdateAvgBalance(i, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvgBalance", reflect.TypeOf((*MockEndOfDayUCItf)(nil).UpdateAvgBalance), i, data)
}

// UpdateBenefit mocks base method.
func (m *MockEndOfDayUCItf) UpdateBenefit(i int, data []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBenefit", i, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBenefit indicates an expected call of UpdateBenefit.
func (mr *MockEndOfDayUCItfMockRecorder) UpdateBenefit(i, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBenefit", reflect.TypeOf((*MockEndOfDayUCItf)(nil).UpdateBenefit), i, data)
}

// UpdateLimitedBalance mocks base method.
func (m *MockEndOfDayUCItf) UpdateLimitedBalance(i int, data []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLimitedBalance", i, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLimitedBalance indicates an expected call of UpdateLimitedBalance.
func (mr *MockEndOfDayUCItfMockRecorder) UpdateLimitedBalance(i, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLimitedBalance", reflect.TypeOf((*MockEndOfDayUCItf)(nil).UpdateLimitedBalance), i, data)
}
