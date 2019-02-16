// Code generated by MockGen. DO NOT EDIT.
// Source: udp-loadbalancer.go

// Package udp_loadbalancer is a generated GoMock package.
package udp_loadbalancer

import (
	gomock "github.com/golang/mock/gomock"
	proto "github.com/k8s-nativelb/pkg/proto"
	reflect "reflect"
)

// MockUdpLoadBalancerInterface is a mock of UdpLoadBalancerInterface interface
type MockUdpLoadBalancerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUdpLoadBalancerInterfaceMockRecorder
}

// MockUdpLoadBalancerInterfaceMockRecorder is the mock recorder for MockUdpLoadBalancerInterface
type MockUdpLoadBalancerInterfaceMockRecorder struct {
	mock *MockUdpLoadBalancerInterface
}

// NewMockUdpLoadBalancerInterface creates a new mock instance
func NewMockUdpLoadBalancerInterface(ctrl *gomock.Controller) *MockUdpLoadBalancerInterface {
	mock := &MockUdpLoadBalancerInterface{ctrl: ctrl}
	mock.recorder = &MockUdpLoadBalancerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUdpLoadBalancerInterface) EXPECT() *MockUdpLoadBalancerInterfaceMockRecorder {
	return m.recorder
}

// UpdateFarm mocks base method
func (m *MockUdpLoadBalancerInterface) UpdateFarm(arg0 *proto.Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFarm", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFarm indicates an expected call of UpdateFarm
func (mr *MockUdpLoadBalancerInterfaceMockRecorder) UpdateFarm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFarm", reflect.TypeOf((*MockUdpLoadBalancerInterface)(nil).UpdateFarm), arg0)
}

// RemoveFarm mocks base method
func (m *MockUdpLoadBalancerInterface) RemoveFarm(arg0 *proto.Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFarm", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFarm indicates an expected call of RemoveFarm
func (mr *MockUdpLoadBalancerInterfaceMockRecorder) RemoveFarm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFarm", reflect.TypeOf((*MockUdpLoadBalancerInterface)(nil).RemoveFarm), arg0)
}

// LoadInitData mocks base method
func (m *MockUdpLoadBalancerInterface) LoadInitData(arg0 *proto.InitAgentData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadInitData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadInitData indicates an expected call of LoadInitData
func (mr *MockUdpLoadBalancerInterfaceMockRecorder) LoadInitData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadInitData", reflect.TypeOf((*MockUdpLoadBalancerInterface)(nil).LoadInitData), arg0)
}

// StartEngine mocks base method
func (m *MockUdpLoadBalancerInterface) StartEngine() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartEngine")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartEngine indicates an expected call of StartEngine
func (mr *MockUdpLoadBalancerInterfaceMockRecorder) StartEngine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartEngine", reflect.TypeOf((*MockUdpLoadBalancerInterface)(nil).StartEngine))
}

// ReloadEngine mocks base method
func (m *MockUdpLoadBalancerInterface) ReloadEngine() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadEngine")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadEngine indicates an expected call of ReloadEngine
func (mr *MockUdpLoadBalancerInterfaceMockRecorder) ReloadEngine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadEngine", reflect.TypeOf((*MockUdpLoadBalancerInterface)(nil).ReloadEngine))
}

// StopEngine mocks base method
func (m *MockUdpLoadBalancerInterface) StopEngine() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopEngine")
}

// StopEngine indicates an expected call of StopEngine
func (mr *MockUdpLoadBalancerInterfaceMockRecorder) StopEngine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopEngine", reflect.TypeOf((*MockUdpLoadBalancerInterface)(nil).StopEngine))
}

// WriteConfig mocks base method
func (m *MockUdpLoadBalancerInterface) WriteConfig() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteConfig")
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteConfig indicates an expected call of WriteConfig
func (mr *MockUdpLoadBalancerInterfaceMockRecorder) WriteConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteConfig", reflect.TypeOf((*MockUdpLoadBalancerInterface)(nil).WriteConfig))
}