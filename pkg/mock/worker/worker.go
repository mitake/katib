// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeflow/katib/pkg/manager/worker (interfaces: Interface)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/kubeflow/katib/pkg/api"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CleanWorkers mocks base method
func (m *MockInterface) CleanWorkers(arg0 string) error {
	ret := m.ctrl.Call(m, "CleanWorkers", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanWorkers indicates an expected call of CleanWorkers
func (mr *MockInterfaceMockRecorder) CleanWorkers(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanWorkers", reflect.TypeOf((*MockInterface)(nil).CleanWorkers), arg0)
}

// IsWorkerComplete mocks base method
func (m *MockInterface) IsWorkerComplete(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "IsWorkerComplete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsWorkerComplete indicates an expected call of IsWorkerComplete
func (mr *MockInterfaceMockRecorder) IsWorkerComplete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWorkerComplete", reflect.TypeOf((*MockInterface)(nil).IsWorkerComplete), arg0)
}

// SpawnWorker mocks base method
func (m *MockInterface) SpawnWorker(arg0 string, arg1 *api.WorkerConfig) error {
	ret := m.ctrl.Call(m, "SpawnWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SpawnWorker indicates an expected call of SpawnWorker
func (mr *MockInterfaceMockRecorder) SpawnWorker(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpawnWorker", reflect.TypeOf((*MockInterface)(nil).SpawnWorker), arg0, arg1)
}

// StopWorkers mocks base method
func (m *MockInterface) StopWorkers(arg0 string, arg1 []string, arg2 bool) error {
	ret := m.ctrl.Call(m, "StopWorkers", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopWorkers indicates an expected call of StopWorkers
func (mr *MockInterfaceMockRecorder) StopWorkers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopWorkers", reflect.TypeOf((*MockInterface)(nil).StopWorkers), arg0, arg1, arg2)
}

// StoreWorkerLog mocks base method
func (m *MockInterface) StoreWorkerLog(arg0 string) error {
	ret := m.ctrl.Call(m, "StoreWorkerLog", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreWorkerLog indicates an expected call of StoreWorkerLog
func (mr *MockInterfaceMockRecorder) StoreWorkerLog(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreWorkerLog", reflect.TypeOf((*MockInterface)(nil).StoreWorkerLog), arg0)
}

// UpdateWorkerStatus mocks base method
func (m *MockInterface) UpdateWorkerStatus(arg0 string) error {
	ret := m.ctrl.Call(m, "UpdateWorkerStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkerStatus indicates an expected call of UpdateWorkerStatus
func (mr *MockInterfaceMockRecorder) UpdateWorkerStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkerStatus", reflect.TypeOf((*MockInterface)(nil).UpdateWorkerStatus), arg0)
}
