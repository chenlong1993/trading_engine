// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/persistence/tradeVariety_repository.go
//
// Generated by this command:
//
//	mockgen -source=./internal/persistence/tradeVariety_repository.go -destination=./mocks//persistence/trade_variety/tradeVariety_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/duolacloud/crud-core/types"
	variety "github.com/yzimhao/trading_engine/v2/internal/models/variety"
	gomock "go.uber.org/mock/gomock"
)

// MockTradeVarietyRepository is a mock of TradeVarietyRepository interface.
type MockTradeVarietyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTradeVarietyRepositoryMockRecorder
	isgomock struct{}
}

// MockTradeVarietyRepositoryMockRecorder is the mock recorder for MockTradeVarietyRepository.
type MockTradeVarietyRepositoryMockRecorder struct {
	mock *MockTradeVarietyRepository
}

// NewMockTradeVarietyRepository creates a new mock instance.
func NewMockTradeVarietyRepository(ctrl *gomock.Controller) *MockTradeVarietyRepository {
	mock := &MockTradeVarietyRepository{ctrl: ctrl}
	mock.recorder = &MockTradeVarietyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTradeVarietyRepository) EXPECT() *MockTradeVarietyRepositoryMockRecorder {
	return m.recorder
}

// Aggregate mocks base method.
func (m *MockTradeVarietyRepository) Aggregate(c context.Context, filter map[string]any, aggregateQuery *types.AggregateQuery) ([]*types.AggregateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Aggregate", c, filter, aggregateQuery)
	ret0, _ := ret[0].([]*types.AggregateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aggregate indicates an expected call of Aggregate.
func (mr *MockTradeVarietyRepositoryMockRecorder) Aggregate(c, filter, aggregateQuery any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockTradeVarietyRepository)(nil).Aggregate), c, filter, aggregateQuery)
}

// Count mocks base method.
func (m *MockTradeVarietyRepository) Count(c context.Context, query *types.PageQuery) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", c, query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockTradeVarietyRepositoryMockRecorder) Count(c, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockTradeVarietyRepository)(nil).Count), c, query)
}

// Create mocks base method.
func (m *MockTradeVarietyRepository) Create(c context.Context, createDTO *variety.CreateTradeVariety, opts ...types.CreateOption) (*variety.TradeVariety, error) {
	m.ctrl.T.Helper()
	varargs := []any{c, createDTO}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*variety.TradeVariety)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTradeVarietyRepositoryMockRecorder) Create(c, createDTO any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, createDTO}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTradeVarietyRepository)(nil).Create), varargs...)
}

// CreateMany mocks base method.
func (m *MockTradeVarietyRepository) CreateMany(c context.Context, items []*variety.CreateTradeVariety, opts ...types.CreateManyOption) ([]*variety.TradeVariety, error) {
	m.ctrl.T.Helper()
	varargs := []any{c, items}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMany", varargs...)
	ret0, _ := ret[0].([]*variety.TradeVariety)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMany indicates an expected call of CreateMany.
func (mr *MockTradeVarietyRepositoryMockRecorder) CreateMany(c, items any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, items}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMany", reflect.TypeOf((*MockTradeVarietyRepository)(nil).CreateMany), varargs...)
}

// CursorQuery mocks base method.
func (m *MockTradeVarietyRepository) CursorQuery(c context.Context, query *types.CursorQuery) ([]*variety.TradeVariety, *types.CursorExtra, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorQuery", c, query)
	ret0, _ := ret[0].([]*variety.TradeVariety)
	ret1, _ := ret[1].(*types.CursorExtra)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CursorQuery indicates an expected call of CursorQuery.
func (mr *MockTradeVarietyRepositoryMockRecorder) CursorQuery(c, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorQuery", reflect.TypeOf((*MockTradeVarietyRepository)(nil).CursorQuery), c, query)
}

// Delete mocks base method.
func (m *MockTradeVarietyRepository) Delete(c context.Context, id types.ID, opts ...types.DeleteOption) error {
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
func (mr *MockTradeVarietyRepositoryMockRecorder) Delete(c, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTradeVarietyRepository)(nil).Delete), varargs...)
}

// FindBySymbol mocks base method.
func (m *MockTradeVarietyRepository) FindBySymbol(ctx context.Context, symbol string) (*variety.TradeVariety, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBySymbol", ctx, symbol)
	ret0, _ := ret[0].(*variety.TradeVariety)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBySymbol indicates an expected call of FindBySymbol.
func (mr *MockTradeVarietyRepositoryMockRecorder) FindBySymbol(ctx, symbol any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBySymbol", reflect.TypeOf((*MockTradeVarietyRepository)(nil).FindBySymbol), ctx, symbol)
}

// Get mocks base method.
func (m *MockTradeVarietyRepository) Get(c context.Context, id types.ID, opts ...types.GetOption) (*variety.TradeVariety, error) {
	m.ctrl.T.Helper()
	varargs := []any{c, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*variety.TradeVariety)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTradeVarietyRepositoryMockRecorder) Get(c, id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTradeVarietyRepository)(nil).Get), varargs...)
}

// Query mocks base method.
func (m *MockTradeVarietyRepository) Query(c context.Context, query *types.PageQuery) ([]*variety.TradeVariety, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", c, query)
	ret0, _ := ret[0].([]*variety.TradeVariety)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockTradeVarietyRepositoryMockRecorder) Query(c, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockTradeVarietyRepository)(nil).Query), c, query)
}

// QueryOne mocks base method.
func (m *MockTradeVarietyRepository) QueryOne(c context.Context, filter map[string]any) (*variety.TradeVariety, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOne", c, filter)
	ret0, _ := ret[0].(*variety.TradeVariety)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOne indicates an expected call of QueryOne.
func (mr *MockTradeVarietyRepositoryMockRecorder) QueryOne(c, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOne", reflect.TypeOf((*MockTradeVarietyRepository)(nil).QueryOne), c, filter)
}

// Update mocks base method.
func (m *MockTradeVarietyRepository) Update(c context.Context, id types.ID, updateDTO *variety.UpdateTradeVariety, opts ...types.UpdateOption) (*variety.TradeVariety, error) {
	m.ctrl.T.Helper()
	varargs := []any{c, id, updateDTO}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*variety.TradeVariety)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTradeVarietyRepositoryMockRecorder) Update(c, id, updateDTO any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{c, id, updateDTO}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTradeVarietyRepository)(nil).Update), varargs...)
}