// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/infrastructure/repository/diner.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	diner "github.com/purplease/golang-integration-sample/pkg/domain/diner"
	repository "github.com/purplease/golang-integration-sample/pkg/infrastructure/repository"
)

// MockDiners is a mock of Diners interface.
type MockDiners struct {
	ctrl     *gomock.Controller
	recorder *MockDinersMockRecorder
}

// MockDinersMockRecorder is the mock recorder for MockDiners.
type MockDinersMockRecorder struct {
	mock *MockDiners
}

// NewMockDiners creates a new mock instance.
func NewMockDiners(ctrl *gomock.Controller) *MockDiners {
	mock := &MockDiners{ctrl: ctrl}
	mock.recorder = &MockDinersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiners) EXPECT() *MockDinersMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDiners) Create(ctx context.Context, newDiner *diner.Diner) (*diner.Diner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, newDiner)
	ret0, _ := ret[0].(*diner.Diner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDinersMockRecorder) Create(ctx, newDiner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDiners)(nil).Create), ctx, newDiner)
}

// Delete mocks base method.
func (m *MockDiners) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDinersMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDiners)(nil).Delete), ctx, id)
}

// GetAll mocks base method.
func (m *MockDiners) GetAll(ctx context.Context, page, limit int64) (*repository.PaginationResultDiner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, page, limit)
	ret0, _ := ret[0].(*repository.PaginationResultDiner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockDinersMockRecorder) GetAll(ctx, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDiners)(nil).GetAll), ctx, page, limit)
}

// GetByID mocks base method.
func (m *MockDiners) GetByID(ctx context.Context, id int64) (*diner.Diner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*diner.Diner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockDinersMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockDiners)(nil).GetByID), ctx, id)
}

// GetTotalCount mocks base method.
func (m *MockDiners) GetTotalCount(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalCount", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalCount indicates an expected call of GetTotalCount.
func (mr *MockDinersMockRecorder) GetTotalCount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalCount", reflect.TypeOf((*MockDiners)(nil).GetTotalCount), ctx)
}