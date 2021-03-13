// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/currencies/repositories/currency.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/trwndh/game-currency/internal/domain/currencies/entity"
	reflect "reflect"
)

// MockCurrency is a mock of Currency interface
type MockCurrency struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyMockRecorder
}

// MockCurrencyMockRecorder is the mock recorder for MockCurrency
type MockCurrencyMockRecorder struct {
	mock *MockCurrency
}

// NewMockCurrency creates a new mock instance
func NewMockCurrency(ctrl *gomock.Controller) *MockCurrency {
	mock := &MockCurrency{ctrl: ctrl}
	mock.recorder = &MockCurrencyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCurrency) EXPECT() *MockCurrencyMockRecorder {
	return m.recorder
}

// CountByName mocks base method
func (m *MockCurrency) CountByName(ctx context.Context, name string) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByName", ctx, name)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByName indicates an expected call of CountByName
func (mr *MockCurrencyMockRecorder) CountByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByName", reflect.TypeOf((*MockCurrency)(nil).CountByName), ctx, name)
}

// Create mocks base method
func (m *MockCurrency) Create(ctx context.Context, params entity.CurrencyDAO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCurrencyMockRecorder) Create(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCurrency)(nil).Create), ctx, params)
}

// Find mocks base method
func (m *MockCurrency) Find(ctx context.Context) ([]entity.CurrencyDAO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx)
	ret0, _ := ret[0].([]entity.CurrencyDAO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockCurrencyMockRecorder) Find(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCurrency)(nil).Find), ctx)
}