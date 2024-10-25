// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/persistence/asset_repository.go
//
// Generated by this command:
//
//	mockgen -source=./internal/persistence/asset_repository.go -destination=./mocks//persistence/asset/asset_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/duolacloud/crud-core/types"
	asset "github.com/yzimhao/trading_engine/v2/internal/models/asset"
	types0 "github.com/yzimhao/trading_engine/v2/internal/models/types"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockAssetRepository is a mock of AssetRepository interface.
type MockAssetRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAssetRepositoryMockRecorder
	isgomock struct{}
}

// MockAssetRepositoryMockRecorder is the mock recorder for MockAssetRepository.
type MockAssetRepositoryMockRecorder struct {
	mock *MockAssetRepository
}

// NewMockAssetRepository creates a new mock instance.
func NewMockAssetRepository(ctrl *gomock.Controller) *MockAssetRepository {
	mock := &MockAssetRepository{ctrl: ctrl}
	mock.recorder = &MockAssetRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssetRepository) EXPECT() *MockAssetRepositoryMockRecorder {
	return m.recorder
}

// Aggregate mocks base method.
func (m *MockAssetRepository) Aggregate(c context.Context, filter map[string]any, aggregateQuery *types.AggregateQuery) ([]*types.AggregateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Aggregate", c, filter, aggregateQuery)
	ret0, _ := ret[0].([]*types.AggregateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aggregate indicates an expected call of Aggregate.
func (mr *MockAssetRepositoryMockRecorder) Aggregate(c, filter, aggregateQuery any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockAssetRepository)(nil).Aggregate), c, filter, aggregateQuery)
}

// Count mocks base method.
func (m *MockAssetRepository) Count(c context.Context, query *types.PageQuery) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", c, query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockAssetRepositoryMockRecorder) Count(c, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockAssetRepository)(nil).Count), c, query)
}

// Create mocks base method.
func (m *MockAssetRepository) Create(c context.Context, createDTO *asset.CreateAsset, opts ...types.CreateOption) (*asset.Asset, error) {
	m.ctrl.T.Helper()
	varargs := []any{c, createDTO}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*asset.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAssetRepositoryMockRecorder) Create(c, createDTO any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, createDTO}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAssetRepository)(nil).Create), varargs...)
}

// CreateMany mocks base method.
func (m *MockAssetRepository) CreateMany(c context.Context, items []*asset.CreateAsset, opts ...types.CreateManyOption) ([]*asset.Asset, error) {
	m.ctrl.T.Helper()
	varargs := []any{c, items}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMany", varargs...)
	ret0, _ := ret[0].([]*asset.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMany indicates an expected call of CreateMany.
func (mr *MockAssetRepositoryMockRecorder) CreateMany(c, items any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, items}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMany", reflect.TypeOf((*MockAssetRepository)(nil).CreateMany), varargs...)
}

// CursorQuery mocks base method.
func (m *MockAssetRepository) CursorQuery(c context.Context, query *types.CursorQuery) ([]*asset.Asset, *types.CursorExtra, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorQuery", c, query)
	ret0, _ := ret[0].([]*asset.Asset)
	ret1, _ := ret[1].(*types.CursorExtra)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CursorQuery indicates an expected call of CursorQuery.
func (mr *MockAssetRepositoryMockRecorder) CursorQuery(c, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorQuery", reflect.TypeOf((*MockAssetRepository)(nil).CursorQuery), c, query)
}

// Delete mocks base method.
func (m *MockAssetRepository) Delete(c context.Context, id types.ID, opts ...types.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{c, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAssetRepositoryMockRecorder) Delete(c, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAssetRepository)(nil).Delete), varargs...)
}

// Despoit mocks base method.
func (m *MockAssetRepository) Despoit(ctx context.Context, transId, userId, symbol string, amount types0.Amount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Despoit", ctx, transId, userId, symbol, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// Despoit indicates an expected call of Despoit.
func (mr *MockAssetRepositoryMockRecorder) Despoit(ctx, transId, userId, symbol, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Despoit", reflect.TypeOf((*MockAssetRepository)(nil).Despoit), ctx, transId, userId, symbol, amount)
}

// Freeze mocks base method.
func (m *MockAssetRepository) Freeze(ctx context.Context, tx *gorm.DB, transId, userId, symbol string, amount types0.Amount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Freeze", ctx, tx, transId, userId, symbol, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// Freeze indicates an expected call of Freeze.
func (mr *MockAssetRepositoryMockRecorder) Freeze(ctx, tx, transId, userId, symbol, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Freeze", reflect.TypeOf((*MockAssetRepository)(nil).Freeze), ctx, tx, transId, userId, symbol, amount)
}

// Get mocks base method.
func (m *MockAssetRepository) Get(c context.Context, id types.ID, opts ...types.GetOption) (*asset.Asset, error) {
	m.ctrl.T.Helper()
	varargs := []any{c, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*asset.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAssetRepositoryMockRecorder) Get(c, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAssetRepository)(nil).Get), varargs...)
}

// Query mocks base method.
func (m *MockAssetRepository) Query(c context.Context, query *types.PageQuery) ([]*asset.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", c, query)
	ret0, _ := ret[0].([]*asset.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockAssetRepositoryMockRecorder) Query(c, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockAssetRepository)(nil).Query), c, query)
}

// QueryOne mocks base method.
func (m *MockAssetRepository) QueryOne(c context.Context, filter map[string]any) (*asset.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOne", c, filter)
	ret0, _ := ret[0].(*asset.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOne indicates an expected call of QueryOne.
func (mr *MockAssetRepositoryMockRecorder) QueryOne(c, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOne", reflect.TypeOf((*MockAssetRepository)(nil).QueryOne), c, filter)
}

// Transfer mocks base method.
func (m *MockAssetRepository) Transfer(ctx context.Context, transId, from, to, symbol string, amount types0.Amount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, transId, from, to, symbol, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transfer indicates an expected call of Transfer.
func (mr *MockAssetRepositoryMockRecorder) Transfer(ctx, transId, from, to, symbol, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockAssetRepository)(nil).Transfer), ctx, transId, from, to, symbol, amount)
}

// UnFreeze mocks base method.
func (m *MockAssetRepository) UnFreeze(ctx context.Context, tx *gorm.DB, transId, userId, symbol string, amount types0.Amount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnFreeze", ctx, tx, transId, userId, symbol, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnFreeze indicates an expected call of UnFreeze.
func (mr *MockAssetRepositoryMockRecorder) UnFreeze(ctx, tx, transId, userId, symbol, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnFreeze", reflect.TypeOf((*MockAssetRepository)(nil).UnFreeze), ctx, tx, transId, userId, symbol, amount)
}

// Update mocks base method.
func (m *MockAssetRepository) Update(c context.Context, id types.ID, updateDTO *asset.UpdateAsset, opts ...types.UpdateOption) (*asset.Asset, error) {
	m.ctrl.T.Helper()
	varargs := []any{c, id, updateDTO}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*asset.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAssetRepositoryMockRecorder) Update(c, id, updateDTO any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, id, updateDTO}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAssetRepository)(nil).Update), varargs...)
}

// Withdraw mocks base method.
func (m *MockAssetRepository) Withdraw(ctx context.Context, transId, userId, symbol string, amount types0.Amount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", ctx, transId, userId, symbol, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockAssetRepositoryMockRecorder) Withdraw(ctx, transId, userId, symbol, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockAssetRepository)(nil).Withdraw), ctx, transId, userId, symbol, amount)
}