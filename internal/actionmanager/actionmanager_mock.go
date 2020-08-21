// Code generated by MockGen. DO NOT EDIT.
// Source: internal/actionmanager/actionmanager.go

// Package actionmanager is a generated GoMock package.
package actionmanager

import (
	gomock "github.com/golang/mock/gomock"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	chainindex "github.com/olympus-protocol/ogen/internal/chainindex"
	state "github.com/olympus-protocol/ogen/internal/state"
	bls_interface "github.com/olympus-protocol/ogen/pkg/bls/interface"
	primitives "github.com/olympus-protocol/ogen/pkg/primitives"
	reflect "reflect"
	time "time"
)

// MockLastActionManager is a mock of LastActionManager interface
type MockLastActionManager struct {
	ctrl     *gomock.Controller
	recorder *MockLastActionManagerMockRecorder
}

// MockLastActionManagerMockRecorder is the mock recorder for MockLastActionManager
type MockLastActionManagerMockRecorder struct {
	mock *MockLastActionManager
}

// NewMockLastActionManager creates a new mock instance
func NewMockLastActionManager(ctrl *gomock.Controller) *MockLastActionManager {
	mock := &MockLastActionManager{ctrl: ctrl}
	mock.recorder = &MockLastActionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLastActionManager) EXPECT() *MockLastActionManagerMockRecorder {
	return m.recorder
}

// NewTip mocks base method
func (m *MockLastActionManager) NewTip(row *chainindex.BlockRow, block *primitives.Block, state state.State, receipts []*primitives.EpochReceipt) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewTip", row, block, state, receipts)
}

// NewTip indicates an expected call of NewTip
func (mr *MockLastActionManagerMockRecorder) NewTip(row, block, state, receipts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTip", reflect.TypeOf((*MockLastActionManager)(nil).NewTip), row, block, state, receipts)
}

// handleStartTopic mocks base method
func (m *MockLastActionManager) handleStartTopic(topic *pubsub.Subscription) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleStartTopic", topic)
}

// handleStartTopic indicates an expected call of handleStartTopic
func (mr *MockLastActionManagerMockRecorder) handleStartTopic(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStartTopic", reflect.TypeOf((*MockLastActionManager)(nil).handleStartTopic), topic)
}

// StartValidator mocks base method
func (m *MockLastActionManager) StartValidator(valPub [48]byte, sign func(*primitives.ValidatorHelloMessage) bls_interface.Signature) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartValidator", valPub, sign)
	ret0, _ := ret[0].(bool)
	return ret0
}

// StartValidator indicates an expected call of StartValidator
func (mr *MockLastActionManagerMockRecorder) StartValidator(valPub, sign interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartValidator", reflect.TypeOf((*MockLastActionManager)(nil).StartValidator), valPub, sign)
}

// ShouldRun mocks base method
func (m *MockLastActionManager) ShouldRun(val [48]byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldRun", val)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldRun indicates an expected call of ShouldRun
func (mr *MockLastActionManagerMockRecorder) ShouldRun(val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldRun", reflect.TypeOf((*MockLastActionManager)(nil).ShouldRun), val)
}

// shouldRun mocks base method
func (m *MockLastActionManager) shouldRun(pubSer [48]byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "shouldRun", pubSer)
	ret0, _ := ret[0].(bool)
	return ret0
}

// shouldRun indicates an expected call of shouldRun
func (mr *MockLastActionManagerMockRecorder) shouldRun(pubSer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "shouldRun", reflect.TypeOf((*MockLastActionManager)(nil).shouldRun), pubSer)
}

// RegisterActionAt mocks base method
func (m *MockLastActionManager) RegisterActionAt(by [48]byte, at time.Time, nonce uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterActionAt", by, at, nonce)
}

// RegisterActionAt indicates an expected call of RegisterActionAt
func (mr *MockLastActionManagerMockRecorder) RegisterActionAt(by, at, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterActionAt", reflect.TypeOf((*MockLastActionManager)(nil).RegisterActionAt), by, at, nonce)
}

// RegisterAction mocks base method
func (m *MockLastActionManager) RegisterAction(by [48]byte, nonce uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterAction", by, nonce)
}

// RegisterAction indicates an expected call of RegisterAction
func (mr *MockLastActionManagerMockRecorder) RegisterAction(by, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAction", reflect.TypeOf((*MockLastActionManager)(nil).RegisterAction), by, nonce)
}

// GetNonce mocks base method
func (m *MockLastActionManager) GetNonce() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetNonce indicates an expected call of GetNonce
func (mr *MockLastActionManagerMockRecorder) GetNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockLastActionManager)(nil).GetNonce))
}
